// 使用闭包打印斐波那契数列
package main

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()

	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
}
