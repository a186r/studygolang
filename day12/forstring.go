// 从字符串生成切片
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c)
	}
}

// 字节数组对比函数,返回两个字节数组字典顺序的整数对比结果
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}

	// 两数组相等
	return 0
}

// 将一个文件加载到内存，然后搜索其中所有数字并返回一个切片
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}
