// 函数作为参数

package main

import "fmt"

func main() {
	callback(2, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

// 将函数作为参数的最好的例子是函数string.IndexFunc()
