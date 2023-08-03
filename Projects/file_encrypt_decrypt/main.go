/* Objective of the Project:
The following tool allows users to encrypt and decrypt a file using a passphrase. */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

// Declare constants
const (
	ivSize  = 12 // Size of the IV block is 12 bytes (96 bits)
	keySize = 16 // The key size is 16 bytes (128 bits)
)

// Function to encrypt a file
func encryptFile(inputFile, outputFile string, passphrase []byte) error {
	// Read the input file
	readFile, readError := os.ReadFile(inputFile)
	if readError != nil {
		fmt.Println("There was an error reading the file.\n", readError)
		return readError
	}

	// Create an AES cipher block using the provided passphrase
	cipherBlock, cipherError := aes.NewCipher(passphrase)
	if cipherError != nil {
		fmt.Println("There was an error creating an AES cipher block using the passphrase.\n", cipherError)
		return cipherError
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, ivSize)
	if _, ivError := rand.Read(iv); ivError != nil {
		fmt.Println("There was an error generating the initialization vector.\n", ivError)
		return ivError
	}

	// Create a new AES cipher
	aesCipher, aesError := cipher.NewGCM(cipherBlock)
	if aesError != nil {
		fmt.Println("There was an error creating an AES cipher.\n", aesError)
		return aesError
	}

	// Encrypt the input data
	encryptedData := aesCipher.Seal(nil, iv, readFile, nil)

	// Write the encrypted data to the output file along with the IV
	output, outputError := os.Create(outputFile)
	if outputError != nil {
		fmt.Println("There was an error writing to the output file.\n", outputError)
		return outputError
	}
	defer output.Close()

	output.Write(iv)
	output.Write(encryptedData)

	fmt.Println("***** File encrypted successfully! *****")
	fmt.Println()

	return nil
}

// Function to decrypt a file
func decryptFile(inputFile, outputFile string, passphrase []byte) error {
	// Read the encrypted file
	readFile, readError := os.ReadFile(inputFile)
	if readError != nil {
		fmt.Println("There was an error reading the file.\n", readError)
		return readError
	}

	// Create an AES cipher block using the provided passphrase
	cipherBlock, cipherError := aes.NewCipher(passphrase)
	if cipherError != nil {
		fmt.Println("There was an error creating an AES cipher block using the passphrase.\n", cipherError)
		return cipherError
	}

	// Extract IV from the input data (first block)
	iv := readFile[:ivSize]
	encryptedData := readFile[ivSize:]

	// Create a new AES cipher
	aesCipher, aesError := cipher.NewGCM(cipherBlock)
	if aesError != nil {
		fmt.Println("There was an error creating an AES cipher.\n", aesError)
		return aesError
	}

	// Decrypt the data
	decryptData, decryptError := aesCipher.Open(nil, iv, encryptedData, nil)
	if decryptError != nil {
		fmt.Println("There was an error decrypting the data.\n", decryptError)
		return decryptError
	}

	// Write the decrypted data to the output file
	output, outputError := os.Create(outputFile)
	if outputError != nil {
		fmt.Println("There was an error writing to the output file.\n", outputError)
		return outputError
	}
	defer output.Close()

	output.Write(decryptData)

	fmt.Println("***** File decrypted successfully! *****")
	fmt.Println()

	return nil
}

// Main function
func main() {
	fmt.Println("***** WELCOME TO AES FILE ENCRYPT / DECRYPT *****")
	fmt.Println()

	var choice int
	var inputFile, outputFileName string

	fmt.Print("What would you like to do?\nEnter 1 for file encryption.\nEnter 2 for file decryption.\n\nI would like to: ")
	fmt.Scan(&choice)

	if choice == 1 {
		// Calling the file encryption function
		fmt.Println("\nStarting file encryption.....")
		fmt.Println()

		fmt.Print("Please provide the name and / or location of the file to be encrypted (with the file extension - e.g., sampleFile.txt): ")
		fmt.Scan(&inputFile)

		fmt.Print("Input file noted!\n\nPlease provide the name and / or location for the encrypted file (with the file extension - e.g., sampleEncryptedFile.txt): ")
		fmt.Scan(&outputFileName)

		fmt.Print("Output file noted!\n\nPlease enter a passphrase to be used for encryption: ")

		passphrase, passphraseError := term.ReadPassword(int(syscall.Stdin))
		if passphraseError != nil {
			fmt.Println("\nThere was an error reading the password.\n", passphraseError)
			return
		}

		fmt.Println()
		fmt.Println()

		// Truncate or pad the passphrase to 16 bytes (128 bits)
		if len(passphrase) < keySize {
			padding := make([]byte, keySize-len(passphrase))
			passphrase = append(passphrase, padding...)
		} else if len(passphrase) > keySize {
			passphrase = passphrase[:keySize]
		} else {
			return
		}

		encryptError := encryptFile(inputFile, outputFileName, passphrase)
		if encryptError != nil {
			fmt.Println("There was an error encrypting the file.\n", encryptError)
			return
		}
	} else if choice == 2 {
		// Calling the file decryption function
		fmt.Println("\nStarting file decryption.....")
		fmt.Println()

		fmt.Print("Please provide the name and / or location of the file to be decrypted (with the file extension - e.g., sampleEncryptedFile.txt): ")
		fmt.Scan(&inputFile)

		fmt.Print("Input file noted!\n\nPlease provide the name and / or location for the decrypted file (with the file extension - e.g., sampleDecryptedFile.txt): ")
		fmt.Scan(&outputFileName)

		fmt.Print("Output file noted!\n\nPlease enter a passphrase to be used for decryption: ")

		passphrase, passphraseError := term.ReadPassword(int(syscall.Stdin))
		if passphraseError != nil {
			fmt.Println("\nThere was an error reading the password.\n", passphraseError)
			return
		}

		fmt.Println()
		fmt.Println()

		// Truncate or pad the passphrase to 16 bytes (128 bits)
		if len(passphrase) < keySize {
			padding := make([]byte, keySize-len(passphrase))
			passphrase = append(passphrase, padding...)
		} else if len(passphrase) > keySize {
			passphrase = passphrase[:keySize]
		} else {
			return
		}

		decryptError := decryptFile(inputFile, outputFileName, passphrase)
		if decryptError != nil {
			fmt.Println("There was an error decrypting the file.\n", decryptError)
			return
		}
	} else {
		fmt.Println("\nERROR! Invalid Choice! Exiting program.....")
		fmt.Println()
	}

	fmt.Println("Thank you for using AES File Encrypt / Decrypt! Have a nice day!")
	fmt.Println()
}
