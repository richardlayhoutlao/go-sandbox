package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount, expectedReturnRate, years float64
	// var years int

	fmt.Println("Investment Amount:")
	fmt.Scan(&investmentAmount)
	
	fmt.Println("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Years:")
	fmt.Scan(&years)


	futureValue := (investmentAmount) * math.Pow((1+expectedReturnRate/100), (years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
