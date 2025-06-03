package main

import "fmt"


func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
    fmt.Scan(&revenue)
	
    fmt.Print("Expenses: ")
    fmt.Scan(&expenses)

    fmt.Print("Tax Rate: ")
    fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
    fmt.Printf("Profit: %.2f\n", profit)
    fmt.Printf("Ratio (EBT/Profit): %.2f\n", ratio)
}