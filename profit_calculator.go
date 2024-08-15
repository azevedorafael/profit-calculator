package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Print("Earnings before tax: ")
	fmt.Println(ebt)

	fmt.Print("Profit: ")
	fmt.Println(profit)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}