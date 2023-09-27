package main

import (
	"fmt"
)

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {

	var total float64 = 0
	singleItem := cartItem{}
	for i := 0; i < len(cart); i++ {
		singleItem = cart[i]
		var itemTotal = float64(singleItem.quantity) * singleItem.price
		total = total + itemTotal
	}
	return total
}

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	fmt.Println(cart)

	result := calculateTotal(cart)
	fmt.Println(result)
}
