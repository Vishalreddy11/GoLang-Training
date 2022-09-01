package functions_and_customers

type Customers struct {
	First_Name, Last_Name string
	Account_Number        int
	Account_Balance       float64
}

//Customers account is initialized with first_name
var (
	Vishal = Customers{"Vishal", "Yellati", 1234, 5000}
	Alex   = Customers{"Alex", "Casey", 2345, 10000}
)

func Addmoney(Account_name *Customers, Amounttoadd float64) {
	Account_name.Account_Balance = Account_name.Account_Balance + Amounttoadd
}

func Withdrawmoney(Account_name *Customers, Amounttowithdraw float64) {
	Account_name.Account_Balance = Account_name.Account_Balance - Amounttowithdraw
}

func Transfermoney(Subtractfromaccount *Customers, amounttotransfer float64, addtoaccount Customers) {
	Subtractfromaccount.Account_Balance = Subtractfromaccount.Account_Balance - amounttotransfer
	addtoaccount.Account_Balance = addtoaccount.Account_Balance + amounttotransfer
}
