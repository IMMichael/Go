package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var colors = [] string{"Red", "Green", "Blue"}//An array with no designated length is effectively a slice.
	fmt.Println(colors)

	colors = append(colors, "Purple") //Adding values using the append function
	fmt.Println(colors)

	colors = append(colors[1:len(colors)]) //Removing the *first* member of the slice using the append function
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])//Removing the *last* member of the slice using the append function
	fmt.Println(colors)

	numbers := make([]int, 3) //Creating an uncapped slice using the make function
	numbers[0] = 1555
	numbers[1] = 12
	numbers[2] = 143
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)


}