package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Println("Enter a string to check if palidrome:")
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i
		if s[i] == s[j] {
			fmt.Println("Is a palidrome")
		} else {
			fmt.Println("Not a palidrome")
		}

	}
}
