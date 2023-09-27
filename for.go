package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	//fmt.Println(colors)
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
}