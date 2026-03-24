package main

import "fmt"

// Constants can be declared outside functions
const LoginToken string = "xyz123" 

func main() {
	// 1. Explicit declaration
	var username string = "hitesh"
	var isLoggedIn bool = true

	// 2. Implicit type inference
	var website = "learncodeonline.in"

	// 3. Walrus operator (Short variable declaration)
	numberOfUsers := 300000 

	fmt.Printf("User %s logged in: %t\n", username, isLoggedIn)
	fmt.Printf("Website: %s, Users: %d\n", website, numberOfUsers)
}