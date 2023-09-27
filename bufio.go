package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		fmt.Print("Enter Text: ")
		input, _ := reader.ReadString('\n')
		fmt.Println("You entered:", input)*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Print a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aFloat)
	}

}
