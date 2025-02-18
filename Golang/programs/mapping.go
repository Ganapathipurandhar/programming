package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	fruits := map[string]float64{
		"apples":  price * 0.9,
		"orange":  price * 1,
		"bananas": price * 0.8,
		"oranges":  price * 1,
	}


	if discount, exists := fruits[product]; exists {
		return discount
	}
	return -1
}

func main() {
	fmt.Println(discountedPrice("apples", 100))  // 90.0
	fmt.Println(discountedPrice("orange", 100))  // 100.0
	fmt.Println(discountedPrice("bananas", 100)) // 80.0
	fmt.Println(discountedPrice("bananas", 100)) // 80.0
	fmt.Println(discountedPrice("oranges", 100)) // 100.0
}
