package main

import (
	"fmt"
	"math/rand"
)

func main() {

	nickels := 0.05
	dimes := 0.1
	quarters := 0.25
	piggyBank := 0.0

	for piggyBank <= 20 {
		num := rand.Intn(3) + 1
		print(num)
		if num == 1 {
			piggyBank += nickels
		} else if num == 2 {
			piggyBank += quarters
		} else if num == 3 {
			piggyBank += dimes
		}
		fmt.Printf("Current Bank: %.8f\n", piggyBank)
	}
}
