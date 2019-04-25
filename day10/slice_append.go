// // 给定一个切片sl， 将一个[]byte数组追加到sl后面。写一个函数Append(slice, data []byte) []byte，该函数在sl不能存储更多数据的时候自动扩容
// package main

// import "fmt"

// func main() {
// 	sl := make([]byte, 5, 20)
// 	b := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'b', 'o', 'p', 'q', 'r', 's'}

// 	// 切片sl信息
// 	fmt.Printf("切片sl的长度是：%d，容量是：%d\n", len(sl), cap(sl))
// 	sl = append(sl, b)
// 	fmt.Println("append之后的slice的长度是：%d", len(slice))

// }

// // func Append(slice, data []byte) []byte {
// // 	// 将data拼接在slice后面
// // 	slice = Append(slice, data)
// // 	fmt.Println("append之后的slice的长度是：%d", len(slice))
// // 	return nil
// // }
package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 4)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
