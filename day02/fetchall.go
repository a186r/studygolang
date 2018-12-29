package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。main函数本身也运行在一个goroutine中。
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go function表示创建一个新的goroutine，并在这个新的goroutine执行这个函数
		// 对每一个命令行参数，都是用go关键字来创建一个goroutine，并且让这个函数异步执行http.Get方法
		go fetch(url, ch)
	}

	// 每当请求返回内容时，fetch函数都会往ch这个channel里写入一个字符串，由这个for循环来处理并打印channel里的字符串
	// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，知道另一个goroutine往这个
	// channel写入或者接收值
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// io.Copy会把相应的Body的内容拷贝到ioutil.Discard输出流中。因为我们需要这个方法返回的字节数，但是又不想要其内容
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
