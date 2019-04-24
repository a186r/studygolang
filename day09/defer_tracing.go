// // package main

// // import "fmt"

// // func main() {
// // 	b()
// // }

// // func trace(s string) {
// // 	fmt.Println("entering:", s)
// // }

// // func untrace(s string) {
// // 	fmt.Println("leaving:", s)
// // }

// // func a() {
// // 	trace("a")
// // 	defer untrace("a")
// // 	fmt.Println("in a")
// // }

// // func b() {
// // 	trace("b")
// // 	defer untrace("b")
// // 	fmt.Println("in b")
// // 	a()
// // }

// // 上面的代码可以实现代码追踪，在进入和离开某个函数时打印相关的消息
// // 下面是更简便的版本

// package main

// import "fmt"

// func main() {
// 	b()
// }

// func trace(s string) string {
// 	fmt.Println("entering:", s)
// 	return s
// }

// func un(s string) {
// 	fmt.Println("leaving", s)
// }

// func a() {
// 	defer un(trace("a"))
// 	fmt.Println("in a")
// }

// func b() {
// 	defer un(trace("b"))
// 	fmt.Println("in b")
// 	a()
// }

// 下面的代码展示了另一种在调试时使用defer的手法
package main

import (
	"io"
	"log"
)

func main() {
	func1("Go")
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()

	return 7, io.EOF
}
J