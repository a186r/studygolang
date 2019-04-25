package main

import "fmt"

func main() {
	// 将Add2函数赋值给p2，并调用
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))

	// 调用Adder
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
