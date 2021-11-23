package main

import "fmt"

func main() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5

	fmt.Printf("%x %x %x\n", red, green, blue)
	fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue)

	var blood uint8 = 255
	red++
	fmt.Println(blood)
	var number int8 = 127
	number++
	fmt.Println(number)
}
