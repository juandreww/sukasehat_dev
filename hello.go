// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	string := "true"

	switch string {
	case "yes", "true", "1":
		fmt.Println("String is true")
	case "no", "false", "0":
		fmt.Println("String is false")
	}

}
