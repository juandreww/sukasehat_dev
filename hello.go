// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"

	newString := strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	cipherLen := len(newString) / len(keyword)
	stringx := strings.Repeat(keyword, cipherLen)

	fmt.Println(newString)
	count := 0
	var result string
	for _, c := range stringx {
		merje := int32(newString[count]) + 65 - c
		if merje <= 65 {
			merje += 25
		}
		newchar := string(merje)
		result = result + newchar
		count++
	}
	fmt.Println(result)

	count = 0
	var result2 string
	for _, c := range stringx {
		merje := int32(result[count]) + c - 65
		if merje >= 90 {
			merje -= 25
		}
		newchar := string(merje)
		result2 = result2 + newchar
		count++
	}

	fmt.Println(result2)
}
