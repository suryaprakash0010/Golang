package main

import (
	"fmt"
	"time"
)

func main() {
	// --- PART 1: TIME (The 01-02-2006 Rule) ---
	fmt.Println("--- Time Study ---")
	
	presentTime := time.Now()
	fmt.Println("Raw Time:", presentTime)

	// CRITICAL: Go does NOT use YYYY-MM-DD. 
	// It uses a specific reference: Mon Jan 2 15:04:05 MST 2006
	// Memory Trick: 01 (Month), 02 (Day), 03 (Hour/PM), 04 (Minute), 05 (Second), 06 (Year)
	fmt.Println("Formatted:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("Custom Date:", createdDate.Format("01-02-2006 Monday"))

	fmt.Println("\n--- Pointer Class ---")

	// --- PART 2: POINTERS (Memory Management) ---
	myNumber := 23

	// '&' provides the memory address (The "Where")
	var ptr = &myNumber 

	fmt.Println("Address of myNumber:", ptr)   // e.g., 0xc000012088
	
	// '*' provides the value at that address (The "What")
	fmt.Println("Value at that address:", *ptr) 

	// If I change the value at the address, the original variable changes!
	*ptr = *ptr + 2
	fmt.Println("New value of myNumber:", myNumber) // Outputs 25
}