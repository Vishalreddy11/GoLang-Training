Package main

import (
	"github.com/gin-gonic/gin"
)

type Customers struct {
	First_Name, Last_Name string
	Account_Number        int
	Account_Balance       float64
}

func Balance(c *gin.Context) {
	Cust := []Customers{
		{First_Name: "Vishal", Last_Name: "Yellati", Account_Number: 1234, Account_Balance: 5000},
		{First_Name: "Alex", Last_Name: "Casey", Account_Number: 2345, Account_Balance: 10000},
		{First_Name: "Sean", Last_Name: "Mitchell", Account_Number: 3456, Account_Balance: 12000},
	}
	c.JSON(200, Cust)
}

func main() {
	r := gin.New()

	r.GET("/All", Balance)

	r.Run()
}
