package main

import "fmt"

func main() {
	// --- ARRAY ---
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Array:", fruitList)

	// --- SLICE ---
	var vegList = []string{"Potato", "Beans", "Mushroom"}
	// Appending to a slice
	vegList = append(vegList, "Onion")
	fmt.Println("Slice:", vegList)

	// Removing a value from a slice based on index (e.g., remove index 1 "Beans")
	indexToRemove := 1
	// We append everything BEFORE the index, and everything AFTER the index
	vegList = append(vegList[:indexToRemove], vegList[indexToRemove+1:]...)
	fmt.Println("Slice after removal:", vegList)
}