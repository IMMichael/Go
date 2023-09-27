package main

import (
	"fmt"
	"strconv"
)

func main() {

	var value1 string = "16.5"
	var value2 string = "55.5"
	aFloat1, err := strconv.ParseFloat(value1, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aFloat1)
	}
	aFloat2, err := strconv.ParseFloat(value2, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aFloat2)
	}
	valueSum := aFloat1 + aFloat2

	fmt.Println(valueSum)
	fmt.Printf("%T", valueSum)
}
