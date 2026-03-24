package main

import (
	"fmt"
	"sort"
)

func main() {
	// --- PART 1: ARRAYS (Fixed Size) ---
	fmt.Println("--- Arrays ---")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach" // Index 2 will be empty string "" by default

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Length is fixed at:", len(fruitList))

	// --- PART 2: SLICES (Dynamic Size) ---
	fmt.Println("\n--- Slices ---")
	var mySlice = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type: %T\n", mySlice) // Notice no number in the bracket []

	// Adding items (Go handles memory expansion automatically)
	mySlice = append(mySlice, "Mango", "Banana")
	fmt.Println("After append:", mySlice)

	// Slicing the slice (ranges)
	// [1:3] means start at index 1 and stop BEFORE index 3
	mySlice = append(mySlice[1:3]) 
	fmt.Println("Sliced (1 to 3):", mySlice)

	// --- PART 3: THE 'MAKE' KEYWORD ---
	highScores := make([]int, 4) // Initial size 4

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // This would CRASH (out of bounds)

	// BUT, append works because it re-allocates memory!
	highScores = append(highScores, 555, 666, 321)
	fmt.Println("New Highscores:", highScores)

	sort.Ints(highScores) // Sorting the slice
	fmt.Println("Sorted Scores:", highScores)

	// --- PART 4: REMOVING AN ITEM (Video 15) ---
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("\nCourses:", courses)
	
	index := 2 // Let's remove "swift"
	
	// '...' is the variadic operator. It "unpacks" the slice.
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removing index 2:", courses)
}