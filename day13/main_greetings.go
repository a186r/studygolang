// 创建一个程序main_greetings.go能够和用户说"Good Day"或者“Good Night"
// 不同的问候应该放到单独的greetings保重，
// 在同一个包中创建一个IsAM函数返回一个布尔值来判断当前时间是AM还是PM，同样创建IsAfternoon和IsEvening函数
// 使用main_greetings作出合适的问候
package main

import (
	"fmt"

	"./greetings/greetings"
)

func main() {
	name := "James"
	fmt.Println(greetings.GoodDay(name))
}
