package main

import "fmt"

func Maps() {
	// Declare and initialize a map
	employeeIDs := map[string]int{
		"Chandler": 1000000,
		"Joey":     1000001,
		"Monica":   1000002,
		"Phoebe":   1000003,
		"Rachael":  1000004,
		"Ross":     1000005,
	}

	// Access values by key
	fmt.Println("Phoebe's Employee ID: ", employeeIDs["Phoebe"])

	// Add or update values
	employeeIDs["Gunther"] = 1000007
	employeeIDs["Chandler"] = 1000006

	// Delete a key-value pair
	delete(employeeIDs, "Gunther")

	fmt.Println()

	// Check if a key exists in the map
	employeeID, exists := employeeIDs["Monica"]
	if exists {
		fmt.Println("Monica's Employee ID: ", employeeID)
	}

	fmt.Println()

	// Iterate over the map
	for name, empID := range employeeIDs {
		fmt.Printf("%s's Employee ID: %d\n", name, empID)
	}
}
