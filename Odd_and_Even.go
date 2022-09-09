package main

import (
	"fmt"
)

var Numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("Odd: ", v)
	}
}

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("Even:", v)
	}
}

func main() {
	Channel_to_odd := make(chan int)
	Channel_to_even := make(chan int)

	go odd(Channel_to_odd)
	go even(Channel_to_even)

	for _, v := range Numbers {
		if v%2 == 0 {
			Channel_to_even <- v
		} else {
			Channel_to_odd <- v
		}
	}
}
