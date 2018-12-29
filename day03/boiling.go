/*
	变量声明，每个源文件以包的声明语句开始，说明该源文件是属于哪个包
	包声明语句之后是import语句导入依赖的其他包，然后是包一级的类型、常量、函数的声明语句
*/
package main

import "fmt"

// 包一级的常量，可以在整个包对应的每个源文件中访问
const boilingF = 212.0

func main() {
	// f和c两个是在main函数内部的变量
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC \n", f, c)
}
