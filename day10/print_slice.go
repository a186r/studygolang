// package main

// import "fmt"

// func main() {
// 	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

// 	slice1 := b[1:4]
// 	slice2 := b[:2]
// 	slice3 := b[2:]
// 	slice4 := b[:]

// 	for _, val := range slice1 {
// 		fmt.Println(val)
// 	}

// 	println(slice1)
// 	println(slice2)
// 	println(slice3)
// 	println(slice4)
// }
package main

import "fmt"

// 计算数组元素和
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr[:]))
}
