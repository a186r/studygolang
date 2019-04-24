// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。

package main

import "fmt"

var input1 int = 18
var input2 int = 22

var rh, rj, rc int

func main() {
	rh, rj, rc = getHJC(input1, input2)
	printValue()

	rh, rj, rc = getHJC2(input1, input2)
	printValue()
}

func printValue() {
	fmt.Printf("两数之和、积、差分别是：%d、%d、%d\n", rh, rj, rc)
}

//使用非命名返回值是一个很糟糕的编程习惯
func getHJC(input1, input2 int) (int, int, int) {
	return input1 + input2, input1 * input2, input1 - input2
}

func getHJC2(input1, input2 int) (H, J, C int) {
	H = input1 + input2
	J = input1 * input2
	C = input1 - input2
	return
}
