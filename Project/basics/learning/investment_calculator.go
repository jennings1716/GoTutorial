package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var rateOfReturn float64 = 10
	var noOfYears float64 = 10
	const inflationRate float64 = 6

	var finalValue = investmentAmount * math.Pow(1+rateOfReturn/100, noOfYears)

	var realValue = finalValue / math.Pow(1+inflationRate/100, float64(noOfYears))

	fmt.Println(finalValue)

	fmt.Println(realValue)

}
