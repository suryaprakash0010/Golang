package main

import "fmt"

// Define a struct (Capital letter means it is public/exported)
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// Creating a struct instance
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 26}
	
	fmt.Println("User details:", hitesh)
	// Accessing fields with dot notation
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}