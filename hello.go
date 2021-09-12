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
	
	m := map[string]int{
		"hello": 5,
		"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok)
		v, ok = m["world"]
		fmt.Println(v, ok)
		v, ok = m["goodbye"]
		fmt.Println(v, ok)
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
