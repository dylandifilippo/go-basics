package main

import (
	"fmt"
	"math"
)

const inflationRate = 3.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Real Future value: %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
