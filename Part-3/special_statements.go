package main

import "fmt"

func add(num1, num2 int) int {
	return num1 + num2
}

func SpecialStatements() {
	// break statement
	fmt.Println("Printing numbers till 13 is met..........")
	for i := 0; i <= 20; i++ {
		if i == 13 {
			break
		}

		fmt.Printf("%d\t", i)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	// return statement
	fmt.Println("Proceeding to add two numbers..........")

	var num1, num2 int
	fmt.Print("Please provide the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Please provide the second number: ")
	fmt.Scan(&num2)

	fmt.Println("The sum of the two numbers are: ", add(num1, num2))

	fmt.Println()
	fmt.Println()

	// continue statement
	fmt.Println("Proceeding to print odd numbers less than 20..........")

	for j := 0; j <= 20; j++ {
		if (j % 2) == 0 {
			continue
		}

		fmt.Printf("%d\t", j)
	}

	fmt.Println()
}
