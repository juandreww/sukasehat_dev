package main

import "fmt"

const (
	idKey = "id"
	nameKey = "name"
)
// const z = 20 * 10
// var h = [12]int{1, 5: 4, 6, 10: 100, 15}
// var j = [...]int{1,2,3,4,5,6,7,8,9,10,11,12} //j == h > true
// var i []int
	
func main() {
	x := make([]int, 0, 5)
x = append(x, 1, 2, 3, 4)
y := x[:2]
z := x[2:]
fmt.Println(cap(x), cap(y), cap(z))
y = append(y, 30, 40, 50)
x = append(x, 60)
z = append(z, 70)
fmt.Println("x:", x)
fmt.Println("y:", y)
fmt.Println("z:", z)
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
