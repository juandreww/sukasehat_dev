package main

import "fmt"

const x int64 = 10
const (
	idKey = "id"
	nameKey = "name"
)
const z = 20 * 10
var h = [12]int{1, 5: 4, 6, 10: 100, 15}
	
func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)
	x = x + 1
	y = "bye"
	fmt.Println(x)
	fmt.Println(y)

	p := 10
	p = 20
	fmt.Println(p)
	p = 30
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
