package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
		var randm = rand.Intn(4)
		fmt.Printf("Facing obstacles at %v\n", randm)
		if randm == 0 {
			fmt.Printf("Program stops at %v\n", randm)
			break
		}
	}
	fmt.Println("Liftoff!")
}
