// // // package main

// // // import "fmt"

// // // const (
// // // 	FIZZ     = 3
// // // 	BUZZ     = 5
// // // 	FIZZBUZZ = 15
// // // )

// // // func main() {
// // // 	for i := 0; i <= 100; i++ {
// // // 		switch {
// // // 		case i%FIZZBUZZ == 0:
// // // 			fmt.Println("FIZZBUZZ")
// // // 		case i%FIZZ == 0:
// // // 			fmt.Println("FIZZ")
// // // 		case i%BUZZ == 0:
// // // 			fmt.Println("BUZZ")
// // // 		default:
// // // 			fmt.Println(i)
// // // 		}
// // // 	}
// // // }
// // package main

// // import "fmt"

// // func main() {
// // 	w, h := 20, 10
// // 	for y := 0; y < h; y++ {
// // 		for x := 0; x < w; x++ {
// // 			fmt.Print("*")
// // 		}
// // 		fmt.Println()
// // 	}
// // }

// package main

// import "fmt"

// func main() {
// 	var i int = 5

// 	for i >= 0 {
// 		fmt.Printf("The variable i is now: %d\n", i)
// 	}
// }
