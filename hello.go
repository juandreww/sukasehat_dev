package main
 
import (
	"bufio"
	"os"
	"fmt"
	"sort"
)
 
func main() {
 
	r := bufio.NewReader(os.Stdin)
 
	var t int
	fmt.Fscan(r, &t)
 
	for i := 0; i < t; i++ {
		var a, b, c, m int
		fmt.Fscan(r, &a, &b, &c, &m)
 
		fmt.Println(solve(a, b, c, m))
	}
}

func solve(a, b, c, m int) string {
	max := a + b + c - 3
 
	var arr []int
	arr = append(arr, a, b, c)
	sort.Ints(arr)
 
	var min int
 //
	if arr[2] > arr[0] + arr[1] + 1 {
		min = arr[2] - (arr[0] + arr[1] + 1)
 
	} else {
		min = 0
	}
	//
	if m < min || m > max {
		return "NO"
	}
 
	return "YES"
}