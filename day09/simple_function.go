package main

import "fmt"

func main() {
	fmt.Printf("2*5*6的结果是：%d\n", Multiply3Nums(2, 5, 6))
}

func Multiply3Nums(a int, b int, c int) int {
	return a * b * c
}
