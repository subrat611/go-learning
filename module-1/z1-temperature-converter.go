package main;

import (
	"fmt"
)

func main() {
	var celsius float64
	var x string
	fmt.Println(x)
	fmt.Print("Enter temperature in celsius: ")
	fmt.Scan(&celsius) // taking the input and put in the address

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
}