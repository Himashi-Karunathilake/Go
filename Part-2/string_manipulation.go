package main

import (
	"fmt"
	"strings"
)

func StringManipulation() {
	// Finding the length of a string
	str1 := "Hello World!"
	fmt.Printf("Length of the string \"%s\": %d\n", str1, len(str1))

	fmt.Println()

	// Counting the occurrences of a string
	str2 := "Hello, Hello, Hello"
	fmt.Printf("Number of occurrences of the word \"Hello\" in the string \"%s\": %d\n", str2, strings.Count(str2, "Hello"))

	fmt.Println()

	// Checking the presence of a string
	fmt.Printf("The word \"World\" exists in the string \"%s\": %t\n", str1, strings.Contains(str1, "World"))

	fmt.Println()

	// Checking the index location of a string
	fmt.Printf("The index location of the word \"World\" in the string \"%s\": %d\n", str1, strings.Index(str1, "World"))
}
