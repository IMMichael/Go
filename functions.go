package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")

	sum := addValues(5,8)
	fmt.Println(sum)

	addAll := addAllValues(3,5,6)
	fmt.Println(addAll)

}

func addValues (value1, value2 int) int{ //if the variables of the function are of the same type it need not be declared again (in this case both are int)
	return value1 + value2
}

func addAllValues (values ... int) int{
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}