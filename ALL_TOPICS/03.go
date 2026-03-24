package main

import "fmt"

func main() {
	myNumber := 23
	
	// ptr is a pointer to an integer memory address
	var ptr *int = &myNumber

	fmt.Println("Memory address of myNumber:", ptr) 
	fmt.Println("Actual value stored at address:", *ptr) 

	// Changing value via pointer
	*ptr = *ptr + 2
	fmt.Println("New value of myNumber:", myNumber) // Outputs 25
}