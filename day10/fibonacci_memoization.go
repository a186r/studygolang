package main

import (
	"fmt"
	"time"
)

const LIM = 1000000

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for index := 0; index < LIM; index++ {
		result = fib(index)
		fmt.Printf("fib(%d) is: %d\n", index, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("最终花费的时间是:%s\n", delta)
}

func fib(n int) (res uint64) {
	// 检查fib(n)是否已经存在
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}

	fibs[n] = res

	return
}
