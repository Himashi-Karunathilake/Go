/* Objective of the Project:
Check if a brute-force or dictionary-based password cracking can be performed on a user provided password.

NOTE: 	The password file that is being used was obtained from: https://github.com/berandal666/Passwords/blob/master/10_million_password_list_top_100000.txt */

package main

import (
	"fmt"
	"os" // Provide buffered I/O operations; previously "io/ioutils" package which is now depreciated
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	fmt.Println("Initialization started..........")

	// Open the file and read the contents into a string
	filePath := "passwords_file.txt"
	readContent, readError := os.ReadFile(filePath)
	if readError != nil {
		fmt.Println("There was an error reading the file.\n", readError)
		return
	}

	// Split the string into a list of passwords
	fileContent := strings.Split(string(readContent), "\n")

	// Create a map where the keys are the passwords and the values are true.
	passwordMap := make(map[string]bool)

	for _, password := range fileContent {
		passwordMap[password] = true
	}

	fmt.Println("Initialization completed!")
	fmt.Println()

	// Obtain the password that needs to be verified from the user
	userPassword := ""

	fmt.Print("Please enter the password that you need to verify: ")

	passwordBytes, passwordError := term.ReadPassword(int(syscall.Stdin))
	if passwordError != nil {
		fmt.Println("\nThere was an error reading the password.\n", passwordError)
		return
	}

	fmt.Println()

	// Convert the password bytes to a string
	userPassword = string(passwordBytes)

	// Start the dictionary-based attack
	fmt.Println("\nStarting dictionary-based attack..........")
	fmt.Println("Please wait while the attack completes..........")

	var exist bool

	for password := range passwordMap {
		if password == userPassword {
			fmt.Println("Dictionary-based attack completed!\n\nFetching your results..........")
			fmt.Println("\nOh no! The password you provided was found! Please consider using another, more secure password in your system.")
			exist = true
			break
		} else {
			exist = false
		}
	}

	if !exist {
		fmt.Println("Dictionary-based attack completed!\n\nFetching your results..........")
		fmt.Println("\nCongratulations! The password you entered was not found. The password currently in use is secure.")
	}

	fmt.Println("\nExiting programme!")
}
