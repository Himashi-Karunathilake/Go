package main

import "fmt"

func Loops() {
	// for loop
	fmt.Println("Printing numbers less than 5..........")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", i)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	// while loop implemented using a for loop
	var num int
	fmt.Print("Please enter a number less than or equal to 10: ")
	fmt.Scan(&num)

	if num <= 10 {
		fmt.Printf("Printing numbers less than %d..........\n", num)

		j := 0

		for j < num {
			fmt.Printf("%d\t", j)
			j++
		}
	} else {
		fmt.Printf("ERROR! The number %d is greater than 10.\n", num)
	}

	fmt.Println()
	fmt.Println()

	// range loop
	names := []string{"Chandler", "Joey", "Monica", "Phoebe", "Rachael", "Ross"}
	fmt.Println("Printing the index number and the name of the student..........")
	for index, value := range names {
		fmt.Println(index, value)
	}
}
