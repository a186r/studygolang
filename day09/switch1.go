package main

import "fmt"

func main() {
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Printf("It's equal to 98\n")
	case 100:
		fmt.Printf("It's equal to 100\n")
	default:
		fmt.Printf("It's not equal to 98 or 100\n")
	}
}
