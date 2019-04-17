package main

import (
	"github.com/davecgh/go-spew/spew"
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	net "github.com/libp2p/go-libp2p-net"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// 区块链是一组区块的集合
var Blockchain []Block

var mutex = &sync.Mutex{}

func main(){

	t:=time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock),""}

	Blockchain = append(Blockchain, genesisBlock)

	// 使用golog记录消息
	golog.SetAllLoggers(gologging.INFO)

	// Parse options form the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	secio := flag.Bool("secio", false, "enable secio")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0{
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Make a host that listens on the given multiaddress
	ha, err := makeBasicHost(*listenF, *secio, *seed)
	if err != nil {
		log.Fatal(err)
	}

	if *target == "" {
		log.Println("listening for connections")

		// Set a stream handler on host A. /p2p/1.0.0 is a user-defined protocol name.
		ha.SetStreamHandler("/p2p/1.0.0", handleStream)

		select{}
	}else {
		ha.SetStreamHandler("/p2p/1.0.0", handleStream)

		ipfsaddr, err := ma.NewMultiaddr(*target)
		if err != nil {
			log.Fatalln(err)
		}

		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			log.Fatalln(err)
		}

		peerid, err := peer.IDB58Decode(pid)
		if err != nil {
			log.Fatalln(err)
		}

		targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", peer.IDB58Decode(peerid)))
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr) 

		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

		log.Println("opening stream")

		s, err:=ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil{
			log.Fatalln(err)
		}

		rw:=bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		go writeData(rw)
		go readData(rw)

		select{}
	}
}

// 编写允许创建主机的方法，当第一个节点运行我们的程序时，它应当充当一个主机，其他节点连接到主机。
func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {

	// 如果种子是0，给定更多的随机性
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	// 为此主机生成密钥对，我们将使用它获取有效的主机标示
	// 生成公私钥，以便主机保持安全。
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	// 这里开始构造其他对等体可以连接的地址
	opts := []libp2p.Option{
		libp2p.ListenAddrString(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
	}

	if !secio {
		opts = append(opts, libp2p.NoEncryption())
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	//构建主机多地址
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID(), Pretty()))

	//现在构建多个地址来访问该主机
	addr := basicHost.Addr()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)
	if secio {
		log.Printf("Now run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \"go run main.go -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

	return basicHost, nil
}

// 现在我们需要允许主机处理传入的数据流，当另一个节点连接到我们主机并想要提交一个新的区块来覆盖我们自己的时候，我们需要逻辑来确定我们是否应该接受它
// 当我们向区块链添加新块时，我们想将它广播给我们相连接的对等体，所以我们也需要逻辑来做到这一点
func handleStream(s net.Stream) {
	log.Println("Got a new stream!")

	//create a buffer stream for non blocking read and write
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)
	go writeData(rw)

	// stream 's' will stay open until you close it (or the other side closes it)
}

// 我们需要先创建我们读取数据的方法
func readData(rw *bufio.ReaderWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if str == "" {
			return
		}

		if str != "\n" {
			chain := make([]Block, 0)
			if err := json.Unmarshal([]bytes(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()

			if len(chain) > len(Blockchain) {
				Blockchain = chain
				bytes, err := json.MarshalIndent(Blockchain, "", " ")
				if err !+ nil {
					log.Fatal(err)
				}

				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}

			mutex.Unlock()
		}
	}
}

// 现在我们已经接受了对等节点的区块链， 如果我们在我们的区块链增加一个新的区块，我们需要一个方法让其他对等节点知道，这样他们就可以接受我们的。我们使用writeData来实现
func writeData(rw *bufio.ReadWriter) {
	
	go func () {
		for {
			time.Sleep(5 * time.Second)
			mutex.Lock()
			bytes, err := json.Marshal(Blockchain)
			if err != nil {
				log.Println(err)
			}
			mutex.Unlock()

			mutex.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			mutex.Unlock()
		}
	}()

	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sendData = strings.Replace(sendData, "\n", "", -1)
		bpm, err := strconv.Atoi(sendData)
		if err != nil{
			log.Fatal(err)
		}

		newBlock := generateBlock(Blockchain[len(Blockchain)-1], bpm)

		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]){
			mutex.Lock()
			Blockchain = append(Blockchain, newBlock)
			mutex.Unlock()
		}

		bytes,err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(Blockchain)

		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()
	}

}

func isBlockValid(newBlock, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
