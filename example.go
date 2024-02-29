package main

import (
	"fmt"
	"math"
)

func main() {
	var year int
	var interestRate float64
	var investmentAmount float64
	var inflationRate = 6.5
	fmt.Print("Input your investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Input the interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Print("input your investment year: ")
	fmt.Scan(&year)

	futureValue := investmentAmount * math.Pow(1+interestRate/100,float64(year)) 
	actualValue := futureValue/math.Pow(1+inflationRate/100,float64(year))
	println(futureValue)
	println(actualValue)
}