package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Println(ms.i1)
	fmt.Println(ms.f1)
	fmt.Println(ms.str)
	fmt.Println(ms)

	// 初始化一个结构体实例的更简短和惯用的方式如下
	ms := &struct1{10, 15.5, "abc"}
}
