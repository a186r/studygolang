// // 写一个新的版本，主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的Fibonacci切片

// package main

// import "fmt"

// func main() {
// 	// 使用new创建切片
// 	// var p *[]int = new([]int) // *p == nil;with len and cap 0
// 	// p := new([]int)

// 	// 使用make创建切片
// 	// p := make([]int, 0)

// 	//但是以上两种方式都不实用，使用下面的方法
// 	// var v []int = make([]int, 10, 50)
// 	// // 或者
// 	// v := make([]int, 10, 50)
// 	// // 这样就分了一个有50个int的数组，并且创建了一个长度为10，容量为50的切片v，该切片指向数组的前10个元素
// 	s := make([]byte, 5)

// 	s = s[2:4]

// 	fmt.Println(len(s))
// 	fmt.Println(cap(s))
// }

// 通过buffer串联字符串
// 这种方式实现比使用+=更节省内存和cpu，尤其是要串联的字符串数目特别多的时候
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")
}
