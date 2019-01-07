/*
让我们来尝试另外一种策略，如果错误的发生是偶然性的，或者由不可预知的问题导致。
一个明智的选择是重新尝试失败的操作，在重试时我们需要限制重试的时间间隔或者次数，防止无限制的重试。
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "www.google.com"
	// 如果错误发生后，程序无法继续运行，我们就可以采取第三种策略：输出错误信息并结束程序
	// 需要注意的是这种策略只应在main中执行。
	// 对于库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down:%v\n", err)
		os.Exit(1)
	}

	// 调用log.Fatalf可以更简洁的代码达到与上面相同的效果，log中的所有函数，都默认会在错误信息之前输出时间信息
	if err := WaitForServer(url); err != nil {
		log.SetPrefix("wait: ")
		log.SetFlags(0)
		log.Fatalf("Site is down:%v\n", err)
	}

	// 第四种策略，有时，我们只需要输出错误信息就足够了，不需要终端程序的运行，我们可以通过log包提供函数
	// if err := Ping(); err != nil {
	// 	log.Printf("ping failed:%v;networking disabled", err)
	// }
}

func WaitForServer(url string) error {
	// 定义超时时间
	const timeout = 1 * time.Minute
	// deadline
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //success
		}
		log.Printf("server not responding (%s);retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
