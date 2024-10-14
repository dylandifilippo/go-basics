package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt, eat, ratio := calculateRevenue(revenue, expenses, taxRate)

	// print ebt
	fmt.Println("Earnings before tax is", ebt)
	// print eat
	fmt.Println("Earnings after tax is", eat)
	// print ratio
	fmt.Println("Ratio is", ratio)
}

func calculateRevenue(revenue, expenses, taxRate float64) (ebt float64, eat float64, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat
	return ebt, eat, ratio
}
