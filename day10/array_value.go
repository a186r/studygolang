// 1.证明当数组赋值时，发生了内存拷贝

// 2. 写一个循环，并用下标给数组赋值，并且将数组打印在屏幕上

package main

import "fmt"

func main() {
	var arr1 [20]int

	for i, _ := range arr1 {
		arr1[i] = i
	}

	fmt.Println(arr1)
}
