package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	var sunshine = "Emerging from the cave, your eyes meet the midday sun"
	var wearShades = strings.Contains(sunshine, "midday sun")
	fmt.Println("You need to wear a shades:", wearShades)

	var walkingInsideTheCave = "Walking inside the cave. You see a beautiful sign. You decided to read"
	var readSign = strings.Contains(walkingInsideTheCave, "read")
	fmt.Println("Your decision is to read the sign:", readSign)
}