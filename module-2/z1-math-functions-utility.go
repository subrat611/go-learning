package main

import (
	"fmt"
)

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func main() {
	fmt.Printf("Add: %.2f\n", add(10.2, 39))
	fmt.Printf("Subtract: %.2f\n", subtract(39.49, 12))
	fmt.Printf("Multiply: %.2f\n", multiply(2.3, 94.3))
}