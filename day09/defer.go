// 关键字defer，类似与Java和C#的finally语句块
// 当有多个defer行为被注册时，它们会以逆序执行(类似栈，后进先出)
package main

import "fmt"

func main() {
	function1()
	f()
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func function1() {
	fmt.Printf("在函数1顶部\n")
	defer function2()
	fmt.Printf("在函数1底部\n")
}

func function2() {
	fmt.Printf("Function2延迟到调用函数结束\n")
}
