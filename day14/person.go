// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type Person struct {
// 	firstName string
// 	lastName  string
// }

// func upPerson(p *Person) {
// 	p.firstName = strings.ToUpper(p.firstName)
// 	p.lastName = strings.ToUpper(p.lastName)
// }

// func main() {
// 	// 1.结构体作为值类型
// 	var pers1 Person
// 	pers1.firstName = "Chris"
// 	pers1.lastName = "Woodward"
// 	upPerson(&pers1)
// 	fmt.Printf("两个名字分别为：%s, %s\n", pers1.firstName, pers1.lastName)

// 	// 2.结构体作为指针
// 	pers2 := new(Person)
// 	pers2.firstName = "Chris"
// 	pers2.lastName = "Woodward"
// 	// 注意，也可以通过解指针的方式来设置值
// 	(*pers2).lastName = "Woodward" //这个好像是指针的那个什么反向什么东西
// 	upPerson(pers2)
// 	fmt.Printf("两个名字分别为：%s,%s\n", pers2.firstName, pers2.lastName)

// 	// 3.结构体字面量初始化
// 	pers3 := &Person{"Chris", "Woodward"}
// 	upPerson(pers3)
// 	fmt.Printf("两个名字分别是:%s,%s\n", pers3.firstName, pers3.lastName)
// }

// go中的类型转换遵循严格的规则，当为结构体定义了一个alias类型时，此结构体类型和它的alias类型都有相同的
// 底层类型，它们可以如示例10.3那样互相转换，同时需要注意其中非法赋值或者转换引起的编译错误
package main

import "fmt"

type number struct {
	f float32
}

type nr number //alias type

func main() {
	a := number{5.0}
	b := nr{5.0}

	var c = number(b)
	fmt.Println(a, b, c)
}
