package main

import "fmt"

// Declare the Expense struct here
// your code goes here
type Expense struct {
	Name   string
	Amount float64
	Date   string
}

// Implement the Total method to calculate the total amount spent
// your code goes here
func Total(expenses []Expense) float64 {
	var total float64
	for _, v := range expenses {
		total += v.Amount
	}

	return total
}
func (expense Expense) getName() string {
	return expense.Name
}

// Implement the getName method on the Expense struct here
// your code goes here

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}
