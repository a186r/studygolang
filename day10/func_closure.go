// package main

// import "fmt"

// func main() {
// 	var f = Adder()

// 	// 可以看到，在调用的过程中，x的变量的值是被保留的
// 	// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
// 	fmt.Print(f(1), " - ")
// 	fmt.Print(f(20), " - ")
// 	fmt.Print(f(100), " - ")
// }

// func Adder() func(int) int {
// 	var x int
// 	return func(delta int) int {
// 		x += delta
// 		return x
// 	}
// }
package main

import "strings"

func main() {
	// 现在可以生成如下函数
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	addPng := MakeAddSuffix(".png")

	// 调用
	addBmp("file")
	addJpeg("file")
	addPng("file")
}

// 工厂函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 判断字符串是否以suffix为后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 可以返回其他函数的函数和接受其他函数作为参数的函数均被称之为高阶函数，是函数式语言的特点
