package main

import "fmt"

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)
}

// 这个方法改变了reply
// reply是一个指向int变量的指针
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
