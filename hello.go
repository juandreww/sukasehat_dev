package main

import (
	"fmt"
)

func main() {
	question := "L fdph, L vdz, L frqtxhuhg"
	for i := 0; i < len(question); i++ {
		c := question[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c > 'Z' {
				c = c - 26
			}
		} else if c >= 'a' && c <= 'z' {
			c = c - 3
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
