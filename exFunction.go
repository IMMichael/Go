package main

import (
	"fmt"
	"strconv"
)

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	var outcomeResult float64 = 0

	number1 := convertInputToValue(input1)
	number2 := convertInputToValue(input2)

	if operation == "+" {
		outcomeResult = addValues(number1, number2)
	}

	if operation == "-" {
		outcomeResult = subtractValues(number1, number2)
	}

	if operation == "*" {
		outcomeResult = multiplyValues(number1, number2)
	}

	if operation == "/" {
		outcomeResult = divideValues(number1, number2)
	}
	return outcomeResult
}

func convertInputToValue(input string) float64 {

	i, _ := strconv.ParseFloat(input, 64)

	return i
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

// Don't touch this section
func main() {
	value1 := "100"
	value2 := "2"
	operation := "/"
	result := calculate(value1, value2, operation)

	fmt.Println(result)
}
