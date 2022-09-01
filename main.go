package main

import (
	"BankingApp/functions_and_customers"
	"fmt"
)

func main() {
	fmt.Println(functions_and_customers.Vishal)
	//to call functions, to add and withdraw enter func_name(Account_name, "&" amount)
	functions_and_customers.Addmoney(&functions_and_customers.Vishal, 100)
	fmt.Println(functions_and_customers.Vishal)
}
