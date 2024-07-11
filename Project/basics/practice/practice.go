package main

import "fmt"

func main() {
	var revenue float64
	var taxRate float64
	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	var earningsBeforeTax = revenue - expenses

	var earningsAfterTax = earningsBeforeTax - ((taxRate / 100) * earningsBeforeTax)

	var ratio = earningsBeforeTax / earningsAfterTax

	fmt.Printf("Earnings Before Tax %.2f \n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax %.2f \n", earningsAfterTax)
	fmt.Printf("ratio %.2f \n", ratio)
}
