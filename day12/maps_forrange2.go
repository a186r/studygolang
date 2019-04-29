// 假如想获取一个map类型的切片，我们必须使用两次make函数，第一次分配切片，第二次分配切片中的每个map元素
package main

import "fmt"

func main() {
	// Version A
	// 类型为map[int]int 的切片，长度为5
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][i] = 2 * i
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}

// 需要注意的是，应该像A中那样通过索引使用切片的map元素，在B版本中获得的项只是map值的一个拷贝而已，所以真正的map元素并没有得到初始化
