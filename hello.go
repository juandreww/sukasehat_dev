package main

import "fmt"
import "math/rand"

const (
	idKey   = "id"
	nameKey = "name"
)

// const z = 20 * 10
// var h = [12]int{1, 5: 4, 6, 10: 100, 15}
// var j = [...]int{1,2,3,4,5,6,7,8,9,10,11,12} //j == h > true
// var i []int

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} // complete loop

	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	} // conditional loop

	for {
		fmt.Println("Hello")
	} // infinite loop

	for {
		// things to do in the loop
		if !CONDITION {
		break
		}
	} // do while loop in golang mode
	
}
//

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
