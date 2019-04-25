// 使用数组计算斐波那契数列数列
package main

import "fmt"

var fibs [1000000]int

func main() {
	fibs[0] = 1
	fibs[1] = 1

	for i := 2; i < len(fibs); i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	for i := 0; i < len(fibs); i++ {
		fmt.Println(fibs[i])
	}
}
