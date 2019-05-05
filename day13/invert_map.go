// 如果map的值类型可以作为key并且所有的value是唯一的，那么通过下面的方法可以简单的做到键值对调
package main

import "fmt"

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value:%v / ", k, v)
	}
}

// 如果原始的value值不唯一，那么这么做肯定会出问题，这种情况下不会报错，但是当遇到不唯一的key时应当直接
// 停止对调，且此时对调后的map很可能没有包含原map的所有键值对，一种解决方法就是仔细检查唯一性并且使用多值
// map，比如使用map[int][]string类型
