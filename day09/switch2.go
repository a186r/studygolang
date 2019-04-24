// package main

// import "fmt"

// func main() {
// 	var num1 int = 7

// 	switch {
// 	case num1 < 0:
// 		fmt.Printf("Number is negative")
// 	case num1 > 0 && num1 < 10:
// 		fmt.Printf("Number is between 0 and 10")
// 	default:
// 		fmt.Printf("Number is 10 or greater\n")
// 	}

// 	//
// 	// switch result := calculate() {
// 	// case result < 0:
// 	// 	..
// 	// case result>0:
// 	// 	...
// 	// }
// }

package main

import "fmt"

func main() {
	fmt.Printf("当前的季节是：%s\n", Season(14))
}

func Season(month int) (name string) {
	switch {
	case month > 1 && month <= 4:
		return "春季"
	case month > 4 && month <= 7:
		return "夏季"
	case month > 7 && month <= 10:
		return "秋季"
	case month > 10 && month <= 12:
		return "冬季"
	default:
		return "md"
	}
}
