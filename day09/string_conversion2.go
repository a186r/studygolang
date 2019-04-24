package main

import (
	"fmt"
	"math"
	"strconv"
)

// func main() {
// 	var orig string = "ABC"
// 	var newS string

// 	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
// 	// 将一个字符串转换为整数
// 	an, err := strconv.Atoi(orig)
// 	// 到这里说明有错误
// 	// if err != nil {
// 	// 	fmt.Printf("orig %s is not an integer - exiting with error \n", orig)
// 	// 	// 遇到错误，直接return提前结束函数的运行
// 	// 	return
// 	// }

// 	// 错误发生的同时终止程序的运行,习惯用法
// 	if err != nil {
// 		fmt.Printf("")
// 		os.Exit(1)
// 	}

// 	fmt.Printf("The integer is %d\n", an)
// 	an = an + 5
// 	newS = strconv.Itoa(an)
// 	fmt.Printf("The new string is: %s\n", newS)

// 	// 习惯用法
// 	// value, err := pack1.Function1(param1)
// 	// if err != nil {
// 	// 	fmt.Printf("")
// 	// 	return err
// 	// }

// 	// 我们尝试通过os.Open方法打开一个名为name的文件
// 	// f, err = os.Open(name)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// doSomething(f)

// 	// 将错误的获取放置在if语句的初始化部分
// 	// if err := file.Chmod(0664); err != nil {
// 	// 	fmt.Printf(err)
// 	// 	return err
// 	// }
// }

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}

func main() {
	f, ok := mySqrt(26.0)
	if ok {
		fmt.Println(f)
	}
}

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

// 实际上， fmt最简单的打印函数也有两个返回值
count, err := fmt.Println(x)