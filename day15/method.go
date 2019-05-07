// // // 结构体上的方法

// // package main

// // import "fmt"

// // type TwoInts struct {
// // 	a int
// // 	b int
// // }

// // func main() {
// // 	two1 := new(TwoInts)
// // 	two1.a = 12
// // 	two1.b = 15

// // 	fmt.Printf("The sum is: %d\n", two1.AddThem())
// // 	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(2))

// // }

// // func (tn *TwoInts) AddThem() int {
// // 	return tn.a + tn.b
// // }

// // func (tn *TwoInts) AddToParam(param int) int {
// // 	return tn.a + tn.b + param
// // }

// // 非结构体类型上的方法
// // package main

// // import "fmt"

// // type IntVector []int

// // func (v IntVector) Sum() (s int) {
// // 	for _, x := range v {
// // 		s += x
// // 	}
// // 	return
// // }

// // func main() {
// // 	fmt.Println(IntVector{1, 2, 3, 4, 5, 6}.Sum())
// // }

// package main

// import "fmt"

// type IntVector []int

// func (v IntVector) Sum() (s int) {
// 	for _, x := range v {
// 		s += x
// 	}
// 	return
// }

// func main() {
// 	fmt.Println(IntVector{1, 2, 3, 4, 5}.Sum())
// }

package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}
