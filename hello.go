package main

import "fmt"

func main() {
	age := 25
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")

	check := age > marsAge
	fmt.Println(check)
}
