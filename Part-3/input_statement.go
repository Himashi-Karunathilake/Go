package main

import "fmt"

func InputStatement() {
	// Declare variable to assign the user input
	var fname string

	fmt.Print("Enter your first name: ")
	// Obtain user input
	fmt.Scan(&fname)

	fmt.Printf("Hi %s!\n", fname)
}
