package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/joho/godotenv"
)

// 代表区块的结构体
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Validator string
}

// 区块链是一系列经过验证的区块的集合
var Blockchain []Block
var tempBlocks []Block

// 候选块
var candidateBlocks = make(chan Block)

// 向所有节点广播获胜的验证者
var announcements = make(chan string)

var mutex = &sync.Mutex{}

// 验证者
var validators = make(map[string]int)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	// 创建创世块
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	httpPort := os.Getenv("PORT")

	// 启动TCP
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :", httpPort)
	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			pickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

// pinckWinner创建一个验证者池，并选择能够模拟一个块到区块链的验证器
// 通过从池中随机选择，按标记的数量加权
func pickWinner() {
	time.Sleep(30 * time.Second)
	mutex.Lock()
	temp := tempBlocks
	mutex.Unlock()

	lotteryPool := []string{}

	if len(temp) > 0 {
		// 略微修改传统的pos算法
		// 从所有提交块的验证者中，根据标记的数量对它们进行加权
		// 再传统的权益证明中，验证者可以参与，但是无需提交待伪造的块

	OUTER:
		for _, block := range temp {
			// 如果一直在验证者池中，则跳过
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// 锁定验证者列表，以防止数据竞争
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()

			k, ok := setValidators[block.Validator]

			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		// 从彩票池中挑选随机赢家
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		// 向区块链添加winner的块，并广播到其他节点
		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()
				for _ = range validators {
					announcements <- "\nWinning calidator:" + lotteryWinner + "\n"
				}

				break
			}
		}
	}

	mutex.Lock()
	tempBlocks = []Block{}
	mutex.Unlock()
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()

	// 验证者地址
	var address string

	// 允许用户分配要下注的令牌数量
	// 令牌数量越多，锻造新块的几率就越大
	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}

		t := time.Now()
		address = calculateHash(t.String())
		validators[address] = balance
		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\nEnter a new BPM:")

	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			for scanBPM.Scan() {
				// 在进行必要的验证后，从stdin中获取BPM并将其添加到区块链
				bpm, err := strconv.Atoi(scanBPM.Text())
				// 如果验证者作恶，直接删除
				if err != nil {
					log.Printf("%v not a number: %var", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}

				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				// 创建新块
				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if isBlockValid(newBlock, oldLastIndex) {
					candidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")
			}
		}
	}()

	// 模拟接收广播
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(Blockchain)
		mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, string(output)+"\n")
	}
}

// 确保区块是正确的，并与前一个区块的hash进行比较
func isBlockValid(newBlock, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// 计算Hash
func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 计算区块Hash
func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(record)
}

// 生成新区块
func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}
