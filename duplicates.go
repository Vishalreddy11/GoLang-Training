// returning duplicates in an array using maps

package duplicates

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 1}
	y := make(map[int]bool)

	for _, value := range x {
		if _, ok := y[value]; !ok {
			y[value] = true
		} else {
			fmt.Println(value)
		}
	}
}
