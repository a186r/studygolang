// 这里从整体角度对程序做个描述
package main

import (
	"fmt"
	"os"
)

func main() {
	// // var声明两个string类型的变量，变量会在声明时直接初始化
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = ""
	// }

	// fmt.Println(s)
	s, sep := "", ""
	// 声明一个变量的几种方式
	// s := ""
	// var s string
	// var s = ""
	// var s string = ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
