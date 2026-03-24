package main

import (
	"bufio"
	"crypto/rand" // For cryptographically secure random numbers
	"fmt"
	"math/big" // Required for crypto/rand.Int
	"os"
	"strconv"
	"strings"
)

func main() {
	// --- PART 1: CONVERSION LOGIC ---
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Why TrimSpace? Because 'input' contains "5\n". ParseFloat would fail on the "\n".
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

	fmt.Println("------------------------------------")

	// --- PART 2: MATH & RANDOMNESS ---
	fmt.Println("Generating a secure random number...")

	// crypto/rand is used for sensitive data (like passwords or tokens)
	// It is much more secure than math/rand which is "pseudo-random"
	
	// We want a random number from 0 to 4 (limit 5)
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	
	fmt.Println("Your secure random number is:", myRandomNum)
}