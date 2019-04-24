package main

import "fmt"

func main() {
	// 可以传任意多的参数进去
	x := min(1, 3, 4, 6, 0)
	fmt.Printf("The min is: %d\n", x)

	// 可以传一个slice进去
	slice := []int{7, 5, 2, 8, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)

	// 换行打印
	PrintV(slice...)
}

// 该函数接受一个变长参数，并对每个元素进行换行打印
func PrintV(s ...int) {
	for i, v := range s {
		fmt.Printf("这是第%d行的参数%d\n", i+1, v)
	}
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
