// // map默认是无序的，如果你想为map排序，需要将key或者value拷贝到一个切片，再对切片排序，然后可以使用
// // 切片的for-range方法，打印出所有的key和value

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// var (
// 	barVal = map[string]int{
// 		"alpha": 34, "bravo": 56, "charlie": 23,
// 		"delta": 87, "echo": 56, "foxtrot": 12,
// 		"golf": 34, "hotel": 16, "indio": 87,
// 		"juliet": 65, "kili": 43, "lima": 98}
// )

// func main() {
// 	fmt.Println("unsorted:")

// 	for k, v := range barVal {
// 		fmt.Printf("Key: %v, Value: %v\n ", k, v)
// 	}

// 	// 创建key的切片
// 	keys := make([]string, len(barVal))
// 	i := 0

// 	// 借用for-range方法，将key存入切片keys中
// 	for k, _ := range barVal {
// 		keys[i] = k
// 		i++
// 	}

// 	// 对存储key的切片进行排序
// 	sort.Strings(keys)
// 	fmt.Println()
// 	fmt.Println("sorted:")
// 	for _, k := range keys {
// 		fmt.Printf("Key: %v, Value: %v\n", k, barVal[k])
// 	}
// }

// // 如果你想要一个排序的列表，最好使用结构体切片，这样会更有效
// type name struct {
// 	key   string
// 	value int
// }

// map默认是无序的，如果想为map排序，需要先将key或者value拷贝到一个切片，再对切片排序，然后可以使用切片的
// for-range方法，打印出所有的key和value

package main

import (
	"fmt"
	"sort"
)

// 先声明一个map
var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98,
	}
)

func main() {
	// 打印出排序之前的map
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value : %v / ", k, barVal[k])
	}

}
