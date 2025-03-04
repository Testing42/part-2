package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected rate of return: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// is used to add a comment in golang
	//fmt.Println("Future Value:", futureValue)
	//the number after . tells you how many decimals you need after the .
	fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for Inflation): %.2f", futureValue, futureRealValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
}
