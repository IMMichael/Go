package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 34, 75
	intSum := i1 + i2 + i3

	fmt.Println("Integer sum is:", intSum)

	f1, f2, f3 := 24.2, 34.5, 56.6
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float sum is now:", floatSum)

	Radius := 15.5
	circumference := Radius * 2 * math.Pi
	fmt.Printf("The circumference is: %.2f", circumference)
}
