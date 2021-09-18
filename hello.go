package main

import "fmt"
// import "math/rand"

const (
	idKey   = "id"
	nameKey = "name"
)

type MyFuncOpts struct {
	FirstName string
	LastName string
	Age int
}
func MyFunc(opts MyFuncOpts) error {

}
	
func main() {
	MyFunc(MyFuncOpts {
		LastName: "Patel",
		Age: 50,
	})
	My Func(MyFuncOpts {
		FirstName: "Joe",
		LastName: "Smith",
	})
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
