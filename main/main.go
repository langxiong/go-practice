package main

import (
	"shopping"
	"fmt"
)

func main() {
	price, exists := shopping.PriceCheck(5)
	fmt.Printf("price %f, %t", price, exists)
}