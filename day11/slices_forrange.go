// package main

// import "fmt"

// func main() {
// 	var slice1 []int = make([]int, 4)

// 	slice1[0] = 1
// 	slice1[1] = 2
// 	slice1[2] = 3
// 	slice1[3] = 4

// 	for ix, val := range slice1 {
// 		fmt.Printf("Slice at %d is: %d\n", ix, val)
// 	}
// }

package main

import "fmt"

func main() {
	// seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	// for ix, val := range seasons {
	// 	fmt.Printf("Season %d is:%s\n", ix, val)
	// }

	// var season string
	// for _, season = range seasons {
	// 	fmt.Printf("%s\n", season)
	// }
	items := [...]int{10, 20, 30, 40, 50}

	for i, item := range items {
		item *= 2
		items[i] = item
	}

	fmt.Println(items)
}
