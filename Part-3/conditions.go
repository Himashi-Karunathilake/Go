package main

import "fmt"

func Conditions() {
	// if statement
	var lname string
	fmt.Print("Please enter your last name (-1 to exit): ")
	fmt.Scan(&lname)

	if lname != "-1" {
		fmt.Printf("%s's Conditions Program!\n", lname)
	}

	fmt.Println()

	// if-else statement
	var age int
	fmt.Print("Please enter your age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("You are a child!")
	}

	fmt.Println()

	// if with initialization
	if num1 := 123480; num1%20 == 0 {
		fmt.Printf("The number %d is divisible by 20.\n", num1)
	}

	fmt.Println()

	// switch statement
	var day string
	fmt.Print("Enter any day of the week: ")
	fmt.Scan(&day)

	switch day {
	case "Monday":
		fmt.Println("It is the first day of the week!")
	case "Saturday", "Sunday":
		fmt.Println("It is a day of the weekend!")
	default:
		fmt.Println("It is a regular weekday!")
	}
}
