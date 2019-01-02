package main

import (
	"fmt"
	"strings"
	"time"
)

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

const IPv4Len = 4

// func con(s string) IP {
// 	var p [IPv4Len]byte
// }

func sss() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	// fmt.Printf()
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// 遍历下标和对应的值
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 遍历值
	for _, v := range a {
		fmt.Println("%d\n", v)
	}

	// 初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
}

const (
	a = 1
	b
	c = 2
	d
)
