package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Weight = 9
	fmt.Printf("%v\n %v\n", poodle.Breed, poodle.Weight)
}

// Creating a new struct named Dog
// - The capital letter makes it exported, so that the rest of the app can read/write to it (same logic applies to the struct's variables)
type Dog struct {
	Breed string
	Weight int
}