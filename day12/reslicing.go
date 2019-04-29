// package main

// import "fmt"

// func main() {
// 	slice1 := make([]int, 0, 10)

// 	// 这里不能使用range，因为range遍历的是切片大小，此时len=0，应该使用正常的for循环遍历cap赋值
// 	// for i := range slice1 {
// 	// 	slice1 = slice1[0 : len(slice1)+1]
// 	// 	slice1[i] = i
// 	// 	fmt.Printf("slice1当前长度是:%d\n", len(slice1))
// 	// }

// 	for i := 0; i < cap(slice1); i++ {
// 		slice1 = slice1[0 : len(slice1)+1]
// 		slice1[i] = i
// 		fmt.Printf("slice当前长度是：%d\n", len(slice1))
// 	}

// 	for i, val := range slice1 {
// 		fmt.Printf("slice1下标为%d的值为%d\n", i, val)
// 	}
// }
