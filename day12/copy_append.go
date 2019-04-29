package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("n切片中的元素个数是:%d\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(make([]int, 3), 4, 5, 6)
	fmt.Println(sl3)

	slbyte := make([]byte, 3, 3)
	slbyte = appendByte(slbyte, 3, 6)
	fmt.Println(slbyte)
	fmt.Printf("此时slbyte的cap是：%d\n", cap(slbyte))
}

// append很好用，但是如果你想掌控整个追加的过程，你可以实现这样一个appendByte方法
func appendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		// 需要扩容
		newSlice := make([]byte, (n+1)*2)
		// 将slice拷贝到newSlice中
		copy(newSlice, slice)
		// 重新赋值
		slice = newSlice
	}

	// 如果slice的容量够用，扩展长度即可
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
