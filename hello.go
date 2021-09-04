package main

import "fmt"

func main() {
	var new int = 1_234
	var x float64 = 10.6
	var y int = new + int(x)
	fmt.Println(y)
}

/*
	Literals:
	// 0d54 is Decimal
	// 0o113 is Octal
	// 0x4a is hexadecimal
	// 40 is int
	// 40u is unsigned int
	// 40l is long
	// 40ul is unsigned long
*/
