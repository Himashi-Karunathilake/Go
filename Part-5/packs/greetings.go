package packs

import "fmt"

// Exported function
func SayHello() {
	fmt.Println("Hello!")
}

// Un-exported function
func sayGoodbye() {
	fmt.Println("Goodbye!")
}
