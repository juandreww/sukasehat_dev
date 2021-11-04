package main

import "fmt"

func main() {
	const lightSpeed = 1680// km/s
	var distance = 96300000// km
	
	var (
		abc = 3
		bcd = 4
	)
	
	var speed, weight = 100, 30
	var hoursPerDay, minutesPerHour = 24, 60
	
	weight -= 2
	
	fmt.Println(distance/lightSpeed/hoursPerDay, "days")
}