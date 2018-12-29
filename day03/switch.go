package main

import "fmt"

func main() {
	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Printf("landed on edge")
	}
}

type Point struct {
	X, Y int
}

var p Point

// go语言提供了指针,&操作符可以返回一个变量的内存地址，并且*操作符可以获取指针指向的变量内容。
// 但是在go语言里没有指针运算，也就是不能像c语言可以对指针进行加减操作。
// 无tag switch
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
