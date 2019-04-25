// 将一个大数组传递给函数会消耗很多内存，有两种方法可以避免这种现象
// 传递数组的指针
// 传递数组的切片

package main

import "fmt"

func main() {
	array := [3]float64{7.0, 4.2, 5.3}
	x := Sum(&array)
	fmt.Printf("数组之和是：%f\n", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
