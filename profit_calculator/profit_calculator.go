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

	// earnings before tax
	ebt := revenue - expenses
	// earnings after tax
	eat := ebt * (1 - taxRate/100)
	// calculate the ratio between earnings after tax and revenue
	ratio := ebt / eat

	// print ebt
	fmt.Println("Earnings before tax is", ebt)
	// print eat
	fmt.Println("Earnings after tax is", eat)
	// print ratio
	fmt.Println("Ratio is", ratio)
}
