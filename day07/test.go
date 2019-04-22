// // // // package main

// // // // import "fmt"

// // // // func main() {
// // // // 	fmt.Println("hello, world")
// // // // 	test()
// // // // }

// // // // func test() {
// // // // 	fmt.Println("World test")

// // // // 	fmt.Println(FunctionName("你好", 2))
// // // // }

// // // // // 一个函数可以有多个返回值
// // // // func FunctionName(a string, b int) (t1 int, t2 string) {
// // // // 	return b, a
// // // // }

// // // package main

// // // import (
// // // 	"os"
// // // )

// // // func main() {
// // // 	const Pi = 3.1415

// // // 	// 类型推断
// // // 	const i int = 5
// // // 	var n string
// // // 	println(n + "5")

// // // 	// 常量用作枚举
// // // 	const (
// // // 		Unknow = 0
// // // 		Female = 1
// // // 		Male   = 2
// // // 	)

// // // 	// iota可以用作枚举值，第一个iota等于0，每当iota在新的一行被使用时，它的值就会自动+1，简写如下：
// // // 	const (
// // // 		a = iota
// // // 		b
// // // 		c
// // // 	)

// // // 	// iota也可以用在表达式中，没遇到一次const关键字，iota就重置为0
// // // 	// time包中一周中每天的名称
// // // 	const (
// // // 		Sunday = iota
// // // 		Monday
// // // 		Tuesday
// // // 		Wednesday
// // // 		Thursday
// // // 		Friday
// // // 		Saturday
// // // 	)

// // // 	// 也可以使用某个类型作为枚举常量的类型
// // // 	type Color int

// // // 	const (
// // // 		RED Color = iota
// // // 		ORANGE
// // // 		YELLOW
// // // 		GREEN
// // // 		BLUE
// // // 		INDIGO
// // // 		VIOLET
// // // 	)

// // // 	// 变量
// // // 	// var aa, ba *int // a,b都是指针类型
// // // 	// var a int
// // // 	// var b bool
// // // 	// var str string

// // // 	// var (
// // // 	// 	a   int
// // // 	// 	b   bool
// // // 	// 	str string
// // // 	// )

// // // 	// 这种因式分解关键字的写法一般用于生命全局变量。
// // // 	// 当一个变量被声明之后，系统自动赋予它该类型的初始值，
// // // 	// var identifier [type] value
// // // 	// var a int =15
// // // 	//
// // // 	var i = 5
// // // 	var b bool = false

// // // 	var str string = "Go says hello to the world"

// // // 	// 不过自动推断类型并不是任何时候都适用的，当你想要给变量的类型并不是自动推断出的某种类型时，你还是需要显式的指定变量的类型
// // // 	var n int64 = 2

// // // 	var (
// // // 		HOME   = os.Getenv("HOME")
// // // 		USER   = os.Getenv("USER")
// // // 		GOROOT = os.Getenv("GOROOT")
// // // 	)

// // // 	// 这种写法主要用于声明包级别的全局变量，当你在函数体内部声明变量时，因该使用简短声明语法
// // // 	aab := 1
// // // 	println(aab)
// // // }

// // package main

// // import (
// // 	"fmt"
// // 	"os"
// // 	"runtime"
// // )

// // func main() {
// // 	// var goos string = runtime.GOOS
// // 	// fmt.Printf("The operation system is :%s\n", goos)
// // 	// path := os.Getenv("PATH")
// // 	// fmt.Printf("Path is %s\n", path)
// // 	var goos string = runtime.GOOS
// // 	fmt.Printf("The operation system is :%s\n", goos)
// // 	path := os.Getenv("PATH")
// // 	fmt.Printf("PATH is %s\n", path)
// // }

// // 值类型和引用类型
// // 值拷贝

// // func Printf(format string, list of valiables to be printed)

// // 简短形式，使用:=赋值操作符

// a := 50
// b := false

// // a和b的类型将由编译器自动推断

// // 字符串类型是utf-8字符的一个序列
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
