package main

import (
	"bufio"
	"fmt"
	"os"
)

// Public constant (Available across packages)
const LoginToken string = "ghabbhhjd" 

func main() {
	// --- PART 1: EXPLORING TYPES ---
	var username string = "hitesh"
	fmt.Printf("Value: %v, Type: %T \n", username, username)

	var isLoggedIn bool = false
	fmt.Printf("Value: %v, Type: %T \n", isLoggedIn, isLoggedIn)

	// uint8 ranges from 0 to 255
	var smallVal uint8 = 255
	fmt.Printf("Value: %v, Type: %T \n", smallVal, smallVal)

	// float64 provides high precision
	var smallFloat float64 = 255.45544511254451885
	fmt.Printf("Value: %v, Type: %T \n", smallFloat, smallFloat)

	// Zero Values: If no value is assigned, Go gives a default
	var anotherVariable int
	fmt.Printf("Default Int Value: %v, Type: %T \n", anotherVariable, anotherVariable)

	// Walrus operator (Short declaration)
	numberOfUser := 300000.0
	fmt.Printf("Value: %v, Type: %T \n", numberOfUser, numberOfUser)

	fmt.Println("------------------------------------")

	// --- PART 2: USER INPUT (COMMA OK SYNTAX) ---
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// We use bufio.NewReader because fmt.Scan stops reading at the first space.
	// os.Stdin tells the reader to look at the Standard Input (your keyboard).
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// Comma OK Syntax: 
	// Go doesn't have try/catch. 'input' stores the data, 
	// and '_' is a blank identifier that would normally store an error.
	input, _ := reader.ReadString('\n') 
	
	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of this rating is %T \n", input)
}