// Creating a Basic caluclator.

package caluclator

import (
	"fmt"
)

func main() {
	var number1 float64
	var number2 float64
	var operator string
	fmt.Print("Enter The First Number: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter The Operator (+, -, *, /, %): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter The Second Number: ")
	fmt.Scanln(&number2)

	switch operator {

	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Cannot be divide by zero")
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1/number2)
		}

	default:
		fmt.Println("Invalid Operator")

	}
}
