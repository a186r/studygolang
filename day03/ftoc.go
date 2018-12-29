package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	var s string
	fmt.Println(s)
	// var f,err = os.Open(name) os.Open返回一个文件或者一个错误
	return (f - 32) * 5 / 9
}

func bianliang() {
	// anim := git.GIF{LoopCount: nframes}
	freq := rand.Float64() * 3.0
	t := 0.0

	i := 100
	var boiling float64 = 100
	var names []string
	// var err error
	var p Point

	i, j := 1, 0
	// :=是一个变量声明语句，=是一个变量赋值操作
	i, j = j, i //交换i和j的值

	f, err := os.Open("")
	if err != nil {
		return err
	}

	f.Close()

	// 这里声明了in和err两个变量
	in, err := os.Open("infile")
	// 这里声明了out和err两个变量，但是err已经在上面进行了声明，所以这里是进行赋值操作
	out, err := os.Create("outfile")

	// 简短变量声明语句中必须要至少声明一个新的变量，下面的代码将不能编译通过
	// out, err := os.Open("xxfile")

	xx := 1
	p := &xx
	fmt.Println(*p)
	*p = 2
	fmt.Println(xx)

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) //true false false
}

//在go语言中，返回函数中局部变量的地址也是安全的，例如下面的代码，调用f时创建局部变量v，在局部变量
// 地址被返回之后依然有效，因为指针p依然引用这个变量
var p = f()

func f() *int {
	v := 1
	// &v表示取变量v的值，这将会产生一个指向该整数变量的指针，指针对应的数据类型是*int
	// 指针被称作指向int类型的指针，如果指针名字为p，那么可以说p指针指向变量x
	// 或者说 p指针保存了x变量的内存地址。同时*p表达式对应p指针指向的变量的值
	return &v
}
