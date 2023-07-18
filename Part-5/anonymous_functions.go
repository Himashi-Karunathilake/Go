package main

import "fmt"

func AnonymousFunctions() {
	// Assign an anonymous function to a variable
	sum := func(num1, num2 int) int {
		return num1 + num2
	}

	// Call the anonymous function
	var num1, num2 int

	fmt.Print("Please provide the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Please provide the second number: ")
	fmt.Scan(&num2)

	fmt.Println()

	fmt.Printf("The addition of the two numbers are: %d\n", sum(num1, num2))
}
