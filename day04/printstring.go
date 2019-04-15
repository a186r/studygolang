package main

import (
	"fmt"
	"sort"
)

func main() {

	// str := "Hello"
	// str2 := "World"

	ages := make(map[string]int)

	ages2 := map[string]int{
		"alice": 31,
		"bob":   26,
	}

	ages["alice"] = 31
	ages["bob"] = 26

	for name, age := range ages2 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	delete(ages2, "alice")
	fmt.Println(ages2["alice"])
	// fmt.Printf("%10s:+SDSF%v= %v \n", str, str2, "世界")

	var names []string

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
