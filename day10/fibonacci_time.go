// 计算程序运行花费的时间
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for index := 0; index <= 60; index++ {
		start1 := time.Now()
		result := fib(index)
		end1 := time.Now()
		fmt.Printf("当前是第%d个数，结果是%d，计算耗时：%s\n", index, result, end1.Sub(start1))
	}

	end := time.Now()

	delta := end.Sub(start)

	fmt.Printf("计算总耗时:%s\n", delta)
}

func fib(n int) (res int) {
	// start := time.Now()
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	// end := time.Now()
	// fmt.Printf("这是第%d个数，耗费时间为：%s\n", n, end.Sub(start))
	return
}
