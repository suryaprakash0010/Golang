package main

import "fmt"

func main() {
	// --- IF / ELSE ---
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Super user"
	} else {
		result = "Exactly 10 logins"
	}
	fmt.Println(result)

	// --- LOOPS (Break & Continue) ---
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {
		if days[i] == "Wednesday" {
			// Skip Wednesday
			continue 
		}
		if days[i] == "Friday" {
			// Stop the loop entirely when reaching Friday
			break 
		}
		fmt.Println(days[i])
	}
}