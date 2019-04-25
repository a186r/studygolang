package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] //不包括5,长度是三，容量是四

	// 给数组赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// 打印切片
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	fmt.Printf("arr1的长度是 %d\n", len(arr1))
	fmt.Printf("slice1的长度是 %d\n", len(slice1))
	fmt.Printf("slice1的容量是 %d\n", cap(slice1))

	// 扩容切片
	slice1 = slice1[0:4]

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	fmt.Printf("slice1的长度是：%d\n", len(slice1))
	fmt.Printf("slice1的容量是：%d\n", cap(slice1))
}
