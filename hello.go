package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var myNumber = 19
	for {
		var randomNumber = rand.Intn(25)
		fmt.Printf("The current random number is: %v\n", randomNumber)
		if randomNumber == myNumber {
			break;
		}
		time.Sleep(time.Second)
	}
	fmt.Println("You have found it.")
}
