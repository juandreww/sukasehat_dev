// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%v : %c \n", c, c)
		z := 'z' + 400
		fmt.Printf("%v : %c \n", z, z)
	}
}
