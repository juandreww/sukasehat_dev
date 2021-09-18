package main

import "fmt"
// import "math/rand"

const (
	idKey   = "id"
	nameKey = "name"
)

// const z = 20 * 10
// var h = [12]int{1, 5: 4, 6, 10: 100, 15}
// var j = [...]int{1,2,3,4,5,6,7,8,9,10,11,12} //j == h > true
// var i []int

func main() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}

	
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
