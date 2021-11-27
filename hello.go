package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792 * 3600
	const secondsPerDay = 86400
	const days = distance / lightSpeed / secondsPerDay
	fmt.Println("Canis Major Dwarf is", days, "light years away.")
}

