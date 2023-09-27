package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var value1 string
	// var value2 string

	fmt.Println("Enter a number:")

	reader1 := bufio.NewReader(os.Stdin)
	numInput, _ := reader1.ReadString('\n')

	aFloat1, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Enter a 2nd number:")
	reader2 := bufio.NewReader(os.Stdin)
	numInput2, _ := reader2.ReadString('\n')

	aFloat2, err := strconv.ParseFloat(strings.TrimSpace(numInput2), 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(aFloat1 + aFloat2)
}
