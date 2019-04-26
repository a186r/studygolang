// package main

// func main() {

// }

// func slicecopy(to, fm slice, width uintptr) {
// 	// 如果源切片或者目标切片有一个长度为0，那就不需要拷贝，直接return
// 	if fm.len == 0 || to.len == 0 {
// 		return 0
// 	}

// 	// n记录下源切片或者目标切片较短的那一个的长度
// 	n := fm.len
// 	if to.len < n {
// 		n = to.len
// 	}

// 	// 如果如参width = 0，也不需要拷贝了，返回较短的切片的长度
// 	if width == 0 {
// 		return n
// 	}
// 	// 如果开启了竞争检测
// 	if raceenabled {

// 	}
// }
// go append实现
package main

import "fmt"

func main() {
	mySlice := make([]byte, 0, 0)
	data := []byte{1, 2, 3}

	mySlice = append(mySlice, data)
	fmt.Printf("slice: %v\n", mySlice)
}

func appendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// func append(slice []byte, data []byte) []byte {
// 	// 判断长度够不够，不够就扩容
// 	lenslice := len(slice)
// 	newslicelen := len(slice) + len(data)

// 	// 如果长度不够，先扩容
// 	if newslicelen > cap(slice) {
// 		// 创建足够长度的新切片
// 		newslice := make([]byte, newslicelen, newslicelen)
// 		// 复制数据到新的切片中
// 		copy(newslice, slice)

// 		slice = newslice
// 	}

// 	for i, val := range data {
// 		slice[lenslice+i] = val
// 	}

// 	return slice
// }
