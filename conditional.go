package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		fmt.Println("Less than zero")
	} else if theAnswer == 0 {
		fmt.Println("Equal to zero")
	} else {
		fmt.Println("Greater than zero")
	}

	fmt.Println(result)

	if x := -42; x < 0 { //Value initialzied as part of the if statement
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

}