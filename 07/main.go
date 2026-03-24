package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("--- Decision Making in Go ---")

	// 1. IF-ELSE LOGIC
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println("Result:", result)

	// --- PRO TIP: Initialization inside IF ---
	// 'num' is declared, assigned, and checked all in one line.
	// Its scope is ONLY inside this if-else block.
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}

	fmt.Println("\n--- Switch Case (Dice Game) ---")

	// 2. SWITCH CASE
	// Seeding is necessary for math/rand to get different results each time
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice is 1: You can open")
	case 2:
		fmt.Println("Move 2 spots")
	case 3:
		fmt.Println("Move 3 spots")
		fallthrough // This forces the execution of the NEXT case as well
	case 4:
		fmt.Println("Move 4 spots (Triggered by 3 or 4)")
		fallthrough
	case 5:
		fmt.Println("Move 5 spots")
	case 6:
		fmt.Println("Move 6 spots and roll again!")
	default:
		fmt.Println("Invalid dice roll!")
	}
}