package main

import "fmt"

func main() {
	fmt.Println(comma("a186ra186ra186ra186r"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// 递归调用自身
	return comma(s[:n-3]) + "," + s[n-3:]
}
