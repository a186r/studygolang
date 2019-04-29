// 给定一个slice s [] int 和一个int类型的长度因子，扩展s使其长度为len(s)*factor
package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}
	sl = setLen(sl, 4)

	fmt.Printf("sl的长度是：%d\n", len(sl))
}

func setLen(s []int, factor int) []int {
	// newS = s[0 : len(s)*factor]
	// 新建扩容切片
	newS := make([]int, len(s)*factor)
	// 拷贝旧的切片到新的切片中
	copy(newS, s)
	s = newS
	return s
}
