package main

import "fmt"

//In Go, only identifiers that start with an uppercase letter are exported and accessible from other packages. Hence, "Slices".
func Slices() {
	// Create a slice
	names := []string{"Anne", "John", "Miley"}

	// Print the second element - "John"
	fmt.Println(names[1])

	// Print new line
	fmt.Println()

	// Print the complete slice
	fmt.Println(names)
}
