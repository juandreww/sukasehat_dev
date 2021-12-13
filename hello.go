// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	cipherLen := len(cipherText) / len(keyword)
	stringx := strings.Repeat(keyword, cipherLen)
	count := 0
	var result string
	for _, c := range stringx {
		merje := int32(cipherText[count]) + c - 65
		if merje > 90 {
			merje -= 25
		}
		newchar := string(merje)
		result = result + newchar
		count++
	}
	fmt.Println(result)
}
