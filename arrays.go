package main

import (
	"fmt"
)

func main (){
	fmt.Println("Arrays")

	var colors [3] string // Here we initialized an array without assigning values for each member
	colors [0] = "Red"    // And now we assign values for each member of the array
	colors [1] = "Green"
	colors [2] = "Blue"

	fmt.Println(colors)    // In this example we used the entire array to be printed
	fmt.Println(colors[1]) // While in this example we used the 2nd member of the array to be printed

	var numbers = [5]int {0,0,2,3,4} // In this example we initialized and assigned the values in one statement
	fmt.Println(numbers)

	fmt.Println("Number of numbers is:", len(numbers)) // Here we used the 'len' function to find the number of members for each array
	fmt.Println("Number of colors is:", len(colors))

}