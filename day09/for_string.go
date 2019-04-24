// package main

// import "fmt"

// func main() {
// 	str := "Go is a beautiful language"

// 	fmt.Printf("The length of str is : %d\n", len(str))

// 	for ix := 0; ix < len(str); ix++ {
// 		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
// 	}

// 	str2 := "日本人"
// 	fmt.Printf("The length of str2 is: %d\n", len(str2))
// 	for ix := 0; ix < len(str2); ix++ {
// 		fmt.Printf("Character on position %d is %c \n", ix, str2[ix])
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	i := 0

// START:
// 	fmt.Printf("The counter is at %d\n", i)
// 	i++
// 	if i < 15 {
// 		goto START
// 	}
// }

package main

func main() {

	// 使用两个嵌套循环
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}

	// 只使用一个循环
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}
