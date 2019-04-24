package main

import "fmt"

// 函数接收两个参数，比较它们的大小，然后从小到大返回这两个数
func main() {
	var min, max int
	min, max = MinMax(2, -1)
	fmt.Printf("Min is :%d, Max is :%d\n", min, max)
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}
