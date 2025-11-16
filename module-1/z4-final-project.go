// CLI calculator project

package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var isStop string
	var operator string

	fmt.Println("Welcome to the Go Calculator!")

	for {
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
	
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		fmt.Println("Enter the operator (+ - * /): ")
		fmt.Scan(&operator)

		switch operator {
			case "+":
				fmt.Printf("Result: %.2f\n", num1+num2)
			case "-":
				fmt.Printf("Result: %.2f\n", num1-num2)
			case "*":
				fmt.Printf("Result: %.2f\n", num1*num2)
			case "/":
				if num2 == 0 {
					fmt.Println("Error: Division by zero!")
				}else {
					fmt.Printf("Result: %.2f\n", num1/num2)
				}
			default:
				fmt.Println("Invalid operator!")
		}

		fmt.Print("Do you want to continue? (y/n): ")
        fmt.Scan(&isStop)

		if isStop != "y" {
            fmt.Println("Goodbye!")
            break
        }
	}

	
}