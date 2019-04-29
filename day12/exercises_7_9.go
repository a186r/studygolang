// 用顺序函数过滤容器，s是前10个整型的切片，构造一个函数Filter，第一个参数是s，第二个参数是一个fn func(int) bool
// 返回满足函数fn的元素切片。通过fn测试方法测试当整型值是偶数时的情况
package main

import "fmt"

func main() {
	// s是前10个整型的切片
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
}

// 构造一个函数Filter，第一个参数是s，第二个参数是一个fn func(int) bool
func Filter(s []int, fn func(int) bool) []int {
	// 返回满足函数fn的元素切片
	var p []int //nil 先创建好一个空切片
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

// even就是要传进去的那个函数，偶数过滤
func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
