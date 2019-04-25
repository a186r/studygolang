package main

import "fmt"

func main() {
	pos := 4
	result, pos := fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is:%d\n", pos, result)

	pos = 10
	result, pos = fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is:%d\n", pos, result)
}

func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		// 下面的写法是错误的,因为函数fibonacci()有多个返回值，所以应该分开接收
		// val, _ = fibonacci(n-1) + fibonacci(n-2)

		// 正确的写法
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}

	pos = n
	return
}
