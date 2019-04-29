package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	res := RemoveStringSlice(slice, 2, 4)
	fmt.Println(res)
}

// 自己写的，这种方法很有瑕疵
// func RemoveStringSlice(slice []int, start, end int) []int {

// 	// 然后新建一个切片，将前面的和后面的都复制进新切片
// 	result := make([]int, len(slice))
// 	copy(result, slice[:start])
// 	copy(result[start:], slice[end:])

// 	return result
// }

// 参考别人的，copy返回的是复制过来的元素的数量
func RemoveStringSlice(slice []int, start, end int) []int {
	result := make([]int, len(slice)-(end-start)-1)
	// 前半段复制到result中，并获取到位置
	at := copy(result, slice[:start])
	// 后半段接着复制进来
	copy(result[at:], slice[end+1:])
	return result
}
