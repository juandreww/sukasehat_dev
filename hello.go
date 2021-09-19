package main

import "fmt"
import "errors"
// import "math/rand"

const (
	idKey   = "id"
	nameKey = "name"
)

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
	out = append(out, base+v)
	}
	return out
}

func divAndRemainder(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
	err = errors.New("cannot divide by zero")
	return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func main() {
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)
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
