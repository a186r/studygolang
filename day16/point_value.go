package main

import "fmt"

type B struct {
	thing int
}

// change()接受一个指向b的指针，并改变它的成员
func (b *B) change() {
	b.thing = 1
}

// write通过拷贝接受B的值并输出B的内容
func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	var b1 B //b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) //b2是指针
	b2.change()
	fmt.Println(b2.write())
}
