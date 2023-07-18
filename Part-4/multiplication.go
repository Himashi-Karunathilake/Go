package main

import "fmt"

func multiply(num1, num2 int) int {
	return num1 * num2
}

func Multiplication() {
	fmt.Println("Performing multiplications..........")
	var num1, num2 int
	fmt.Print("Please provide the first number for the multiplication: ")
	fmt.Scan(&num1)
	fmt.Print("Please provide the second number for the multiplication: ")
	fmt.Scan(&num2)

	// Call the multiply() function
	fmt.Println("The multiplication of the two numbers is: ", multiply(num1, num2))
}
