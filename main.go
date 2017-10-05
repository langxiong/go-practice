package main

import (
	"fmt"
)

func main() {
	power := 1000

	_, power = "xiongl", getPower()
	fmt.Printf("xiongl's power over %d", power)
}

func getPower() int {
	return 9001
}