package main

import "fmt"

func FormatSpecifiers() {
	name := "Rachael"
	age := 24
	height := 1.5

	// Printf formats according to format specifiers
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
