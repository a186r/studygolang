package main

import (
	"fmt"

	"./person"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("Alina")
	fmt.Println(p.FirstName())
}
