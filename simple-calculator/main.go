package main

import (
	"fmt"
	"strconv"
)

func convertToFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

func calculate(value1 string, value2 string) float64 {
	// Convert the first string to a float64
	float1 := convertToFloat(value1)

	// Convert the second string to a float64
	float2 := convertToFloat(value2)
	
	// Calculate and return the result
	return float1 + float2
}

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println(value1, "+", value2, "=", result)
}