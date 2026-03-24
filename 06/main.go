package main

import "fmt"

// Structs are usually defined outside the main function so they are globally accessible
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// --- PART 1: MAPS (Video 16) ---
	fmt.Println("--- Maps in Golang ---")

	// Initializing a map: make(map[KeyType]ValueType)
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS stands for:", languages["JS"])

	// Deleting a key
	delete(languages, "RB")
	fmt.Println("After deleting Ruby:", languages)

	// Iterating over a map (using range)
	// 'key' gets the shorthand, 'value' gets the full name
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	fmt.Println("\n--- Structs in Golang ---")

	// --- PART 2: STRUCTS (Video 17) ---
	// Go does not have Classes or Inheritance. We use Structs for data.
	
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 26}
	
	// Basic print
	fmt.Println(hitesh)
	
	// Detailed print (%+v shows field names)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	
	// Accessing specific fields
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
}