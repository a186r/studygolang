// package main

// import (
// 	"fmt"
// )

// func main() {
// 	bool1 := true

// 	// 注意，这里不要使用if bool1 == true来判断，因为bool1已经是一个bool值
// 	if bool1 {
// 		fmt.Printf("The value is true \n")
// 	} else {
// 		fmt.Printf("The value is false \n")
// 	}

// 	// 这里有一些有用的例子：
// 	// 判断一个字符串是否为空
// 	str := ""

// 	if str == "" {
// 		fmt.Printf("str is null \n")
// 	}

// 	if len(str) == 0 {
// 		fmt.Printf("str is null \n")
// 	}

// 	// fmt.Printf("str is not null \n")

// 	// 判断go程序的操作系统类型，这里可以通过常量runtime.GOOS来判断
// 	// if runtime.GOOS == "windows" {

// 	// } else {

// 	// }
// 	// 这段代码一般本放在init()中执行
// 	var prompt = "Enter a xxxx,s"
// 	func init () {
// 		if runtime.GOOS == "windows"{

// 		}else{

// 		}
// 	}

// 	// 函数Abs()返回一个整型数字的绝对值
// 	func Abs(x int) int {
// 		if x < 0 {
// 			return -x
// 		}else{
// 			return x
// 		}
// 	}

// 	// isGreater用于比较两个整型数字的大小
// 	func isGreater(x,y int) bool{
// 		if x > y{
// 			return true
// 		}else{
// 			return false
// 		}
// 	}
// }
