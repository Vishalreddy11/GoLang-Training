//To remove vowels from a string
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("Enter a string to remove vowels from it:")
	fmt.Scan(&x)
	for i := 0; i < len(x); i++ {
		a := string(x[i])
		if a == "a" || a == "e" || a == "i" || a == "o" || a == "u" {
			x = strings.ReplaceAll(x, a, "")
			i--
		}
	}
	fmt.Println(x)
}
