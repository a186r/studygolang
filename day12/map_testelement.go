package main

import "fmt"

func main() {
	// 判断map中某个值是否存在
	var value int
	var isPresent bool

	map1 := make(map[string]int, 10)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("Beijing是存在的，对应的值是:%d\n", value)
	} else {
		fmt.Printf("Beijing是不存在的")
	}

	value, isPresent = map1["Paris"]
	if isPresent {
		fmt.Println("Paris是存在的")
	} else {
		fmt.Println("Paris是不存在的")
	}

	// 删除一个条目
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]

	if isPresent {
		fmt.Println("Washington是存在的")
	} else {
		fmt.Println("Washington是不存在的")
	}
}
