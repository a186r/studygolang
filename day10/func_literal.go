// 匿名函数
// 将匿名函数赋值给变量并对其进行调用

package main

import "fmt"

func main() {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		// 这里g变量代表的是函数，func(i int)，值是一个内存地址
		// 所以我们实际上拥有的是一个函数值：匿名函数可以被赋值给变量并作为值使用
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
