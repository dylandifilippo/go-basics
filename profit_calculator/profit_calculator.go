package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

	writeProfitToFile(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	for {
		var userInput float64
		fmt.Print(infoText)
		fmt.Scan(&userInput)

		if userInput <= 0 || userInput == 0 {
			fmt.Println("Please enter a positive number")
		} else {
			return userInput
		}
	}
}

func writeProfitToFile(ebt, profit, ratio float64) {
	profitText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRation: %.1f\n", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(profitText), 0644)
}
