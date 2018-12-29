package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

func variable() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

//下面的两个newInt函数有着相同的行为：
// func newInt() *int {
// 	return new(int)
// }

// fun newInt() *int{
// 	var dummy int
// 	return dummy
// }
// new函数类似是一种语法糖，而不是一个新的基础概念

// 每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：
func fornew() {
	p2 := new(int)
	q2 := new(int)
	fmt.Println(p2 == q2) // false
}

var global *int

func f() {
	var x int
	// 必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，，虽然它是在函数内部定义的。
	// 用go语言的术语说，这个x局部变量从函数f中逃逸了。
	x = 1
	global = &x
}

// 相反，当g函数返回时，变量*y是不可达的，也就是说可以马上被回收的。
// 因此*y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储空间，虽然这里用的是new方式。
// 其实在任何时候，并不需要为了编写正确的代码而要考虑变量的逃逸行为，要记住的是，逃逸的变量需要额外分配内存
// 同时对性能的优化可能会产生细微的影响。
func g() {
	y := new(int)
	*y = 1
}

// 赋值
func fz() {
	x = 1                       //命名变量的赋值
	*p = true                   //通过指针间接赋值
	person.name = "bob"         //结构体字段赋值
	count[x] = count[x] * scale // 数组、slice、map的元素赋值
	// count[x] *= scale 上一步的简洁写法，这样可以省去对变量表达式的重复计算
	// 数值变量也可以支持++和--语句（自增和自减是语句而不是表达式，因此x = i++之类的表达式是错误的）
	v := 1
	v++
	v--
}

// 元组赋值
func yz() {
	// 赋值语句右边的表达式将会先进性求值，然后再统一根性左边对应变量的值，
	// 这对于处理又写同时出现在元组赋值语句左右两边的变量很有帮助。
	// 我们可以这样交换两个变量的值
	x, y = y, x
	a[i], a[j] = a[j], a[i]
}

// 或者是计算两个整数值的最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 或者是计算斐波那契数列的第N个数：
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	// 元组赋值也可以使一系列琐碎赋值更加紧凑
	i, j, k = 2, 3, 5
	return x
}
