package main

import (
	"fmt"
)

func main() {
	var fruit = [] string{"Orange", "Banana", "Apple", "Kiwi"}
	var price float64
	price = float64(100 / len(fruit))

	priceList := make(map[string] float64)

	
	for s, data := range fruit {
		priceList[s] = data
		fmt.Println(data)
	}

	
}