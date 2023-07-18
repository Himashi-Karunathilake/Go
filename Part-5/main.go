// The fifth program in Go.

package main

import (
	"fmt"
	"packs"
)

func main() {
	fmt.Println()

	fmt.Println("*************** RUNNING THE GREETINGS.GO FILE ***************")
	packs.SayHello()
	fmt.Println("__________________________________________________________________________________________")
	fmt.Println()

	fmt.Println("*************** RUNNING THE ANONYMOUS_FUNCTIONS.GO FILE ***************")
	AnonymousFunctions()
	fmt.Println("__________________________________________________________________________________________")
	fmt.Println()
}
