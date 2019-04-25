// package main

// import "fmt"

// func main() {
// 	// var arrAge = [5]int{18, 22, 23, 54, 3}
// 	// var arrLazy = [...]int{11, 2, 2, 2, 3}
// 	// var arrLazy = []int{1, 2, 3, 45, 3, 2, 23}
// 	var addKeyValue = [5]string{3: "Alina", 2: "Sam"}
// 	// var arrKeyValue = []string{3: "Alina", 2: "Sam"}

// 	for i := 0; i < len(addKeyValue); i++ {
// 		fmt.Printf("%s\n", addKeyValue[i])
// 	}
// }

// 你可以取任意数组常量的地址来作为指向新实例的指针

package main

import "fmt"

func fp(a *[3]int) { fmt.Println(a) }

func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}
