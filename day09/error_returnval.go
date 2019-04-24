// 编写一个名字为MySqrt的函数，计算一个float64类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。
// 编写两个版本，一个是非命名返回值，一个是命名返回值

// 求平方根的方法是哪个来着

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// 先来试试负数
	fmt.Print("给个负数搞搞")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error!,返回值是:", ret1, err1)
	} else {
		fmt.Println("Ok! 返回值是：", ret1)
	}

	// 再来个正整数
	fmt.Print("试试正整数")
	if result, err2 := MySqrt(23); err2 != nil {
		fmt.Println("Error! 返回值是：", result, err2)
	} else {
		fmt.Println("Ok! 返回值是：", result, err2)
	}

	// 最后是命名返回值
	fmt.Print(MySqrt2(26))
}

func MySqrt(input float64) (float64, error) {
	if input < 0 {
		return float64(math.NaN()), errors.New("负数木得平方根")
	}
	return math.Sqrt(input), nil
}

func MySqrt2(input float64) (result float64, err error) {
	if input < 0 {
		result = float64(math.NaN())
		err = errors.New("负数木得平方根")
	} else {
		result = math.Sqrt(input)
		// err不用赋值，所以在这里它会是默认值nil
	}
	return
}
