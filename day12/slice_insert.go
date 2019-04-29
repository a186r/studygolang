package main

import "fmt"

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0)
	fmt.Println(res)

	res = InsertStringSlice(s, in, 3)
	fmt.Println(res)
}

// func InsertStringSlice(slice, insertion []string, index int) []string {
// 	result := make([]string, len(slice)+len(insertion))
// 	at := copy(result, slice[:index])
// 	at = copy(result[at:], insertion)
// 	copy(result[at:], slice[index:])
// 	return result
// }

func InsertStringSlice(slice, insertion []string，index int) []string {
	// 先建立一个空的，长度是
	result := make([string, len(slice)+len(insertion)])
	// 再建立一个空切片，将slice最后index个数拷贝到result中去
	at := copy(result,slice[:index])
	// 然后将insertion拷贝到后面
	at += copy(result[at:],insertion)
	// 最后将
	copy(result[at:],slice[index:])
	return result
}
