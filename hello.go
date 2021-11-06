package main

import (
	"fmt"
)

func main() {
	var destination = "Mountain"
	
	if destination == "Mountain" {
		fmt.Printf("We are at the Mountain\n")
	} else if destination == "Cave" {
		fmt.Printf("We are at the Cave\n")
	} else {
		fmt.Printf("We are at the Entrance\n")
	}
}