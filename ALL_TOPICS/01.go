package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Golang!")

	// Comma OK Syntax Example
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a rating: ")

	// ReadString returns the input (input) and an error (err).
	// If we don't care about the error, we could use the blank identifier `_`
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
	} else {
		fmt.Println("Thanks for rating:", input)
	}
}