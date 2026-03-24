package main

import "fmt"

func main() {
	// make(map[KeyType]ValueType)
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)

	// Deleting a value
	delete(languages, "RB")
	
	// Looping through a map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}