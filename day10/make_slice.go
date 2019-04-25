package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 10)

	// 给切片赋值
	for index := 0; index < len(slice1); index++ {
		slice1[index] = index * 5
	}

	// 打印切片
	for index := 0; index < len(slice1); index++ {
		fmt.Println(slice1[index])
	}

	fmt.Println("切片的长度是：%d", len(slice1))
	fmt.Println("切片的容量是：%d", cap(slice1))
}
