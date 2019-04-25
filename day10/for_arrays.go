// 数组元素可以通过索引来读取或者修改。

package main

import "fmt"

func main() {
	// 数组的声明格式
	// var identifier [len]type
	// 例如 var arr1 [5]int
	// var arr1 [5]int
	// for index := 0; index < len(arr1); index++ {
	// 	arr1[index] = index + 2
	// }

	// // for index := 0; index < len(arr1); index++ {
	// // 	fmt.Println(arr1[index])
	// // }
	// for i, _ := range arr1 {
	// 	fmt.Println(arr1[i])
	// }
	a := [...]string{"a", "b", "c", "d"}

	for i := range a {
		fmt.Println(a[i])
	}
}
