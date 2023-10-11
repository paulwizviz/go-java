package main

import (
	"fmt"
)

func calculateTotalPrice(basket []string) int {
	prices := map[string]int{
		"Apple":  35,
		"Banana": 20,
		"Melon":  50,
		"Lime":   15,
	}

	itemCount := make(map[string]int)
	totalPrice := 0

	// Count the number of each item in the basket
	for _, item := range basket {
		itemCount[item]++
	}

	// Calculate the total price based on the item counts and pricing rules
	for item, count := range itemCount {
		price := prices[item]

		switch item {
		case "Melon":
			// Apply 'buy one get one free' for Melons
			price = (count/2 + count%2) * price
		case "Lime":
			// Apply 'three for the price of two' for Limes
			price = (count/3*2 + count%3) * price
		default:
			price = count * price
		}

		totalPrice += price
	}

	return totalPrice
}

func main() {
	shoppingBasket := []string{"Apple", "Apple", "Banana", "Melon", "Melon", "Lime", "Lime", "Lime"}
	totalPrice := calculateTotalPrice(shoppingBasket)
	fmt.Printf("Total cost of the shopping basket: %dp\n", totalPrice)
}
