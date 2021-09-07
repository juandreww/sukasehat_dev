package main

import "fmt"

const (
	idKey = "id"
	nameKey = "name"
)
const z = 20 * 10
var h = [12]int{1, 5: 4, 6, 10: 100, 15}
var j = [...]int{1,2,3,4,5,6,7,8,9,10,11,12} //j == h > true
var i []int
	
func main() {
	const y = "hello"

	p := 10
	p = 20
	fmt.Println(p)
	p = 30
	fmt.Println(i, len(i), cap(i))
	i = append(i, 30)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 40)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 50)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 60)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 70)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 80)
	fmt.Println(i, len(i), cap(i))
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
