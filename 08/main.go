package main

import "fmt"

func main() {
	fmt.Println("--- Loops and Control Flow ---")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// 1. Standard For Loop
	for d := 0; d < len(days); d++ {
		fmt.Println("Standard loop:", days[d])
	}

	// 2. Range Loop (Most common for Slices/Maps)
	// 'i' is the index, 'day' is the value
	for i, day := range days {
		fmt.Printf("Index is %v and value is %v\n", i, day)
	}

	fmt.Println("\n--- Loop Logic (Break, Continue, Goto) ---")

	rougueValue := 1

	// 3. 'While' style loop (Go uses 'for' without initialization)
	for rougueValue < 10 {

		if rougueValue == 2 {
			// GOTO: Jumps execution to a specific label
			goto lco 
		}

		if rougueValue == 5 {
			// CONTINUE: Skips the rest of this iteration and goes to the next
			// NOTE: Must manually increment before continue in this style loop
			rougueValue++ 
			continue
		}

		if rougueValue == 8 {
			// BREAK: Stops the loop entirely
			break 
		}

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	// This is a Label (Used by goto)
lco:
	fmt.Println("Jumping at LearnCodeonline.in via GOTO")
}