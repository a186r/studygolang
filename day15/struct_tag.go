// 带标签的结构体
// 结构体中的字段除了有名字和类型之外，还可以有一个可选的标签，它是一个附属字段的字符串
// 可以是文档或者是其他重要标记，标签的内容一般不在编程中使用，只有包reflect能获取它
//它可以在运行时自省类型、属性和方法，比如：在一个变量上调用reflect.TypeOf()可以获取
// 变量的正确类型，如果变量是一个结构体类型，就可以通过Field来索引结构体的字段，然后使用Tag属性
package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	// 获取变量的类型
	ttType := reflect.TypeOf(tt)
	// 使用Field方法索引结构体的字段
	ixField := ttType.Field(ix)
	// 打印Tag属性
	fmt.Printf("%v\n", ixField.Tag)
}
