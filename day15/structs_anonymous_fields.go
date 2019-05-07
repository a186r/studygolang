// 匿名字段和内嵌结构体
package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    //匿名字段
	innerS //匿名字段
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 12.9
	outer.int = 60
	outer.in1 = 6
	outer.in2 = 9

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 50, innerS{5, 9}}
	fmt.Println("outer2 is:", outer2)
}

// 通过类型outer.int的名字来获取存储在匿名字段中的数据，于是可以得出一个结论，在一个结构体中对于
// 每一种数据类型只能有一个匿名字段
