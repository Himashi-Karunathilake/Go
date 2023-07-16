package main

import (
	"fmt"
	"math"
)

func Operators() {
	// Arithmetic Operators
	var1 := 5
	var2 := 2

	fmt.Println("var1 = ", var1, "\tvar2 = ", var2)
	fmt.Println("Addition of the two variables: ", (var1 + var2))
	fmt.Println("Subtraction of the two variables: ", (var1 - var2))
	fmt.Println("Multiplication of the two variables: ", (var1 * var2))
	fmt.Println("Division of var1 by var2: ", (var1 / var2))
	fmt.Println("Remainder when var1 is divided by var2: ", (var1 % var2))

	fmt.Println()

	// Power operator
	fmt.Println("var1 to the power of var2: ", math.Pow(float64(var1), float64(var2))) // Convert the two integers to float64
}
