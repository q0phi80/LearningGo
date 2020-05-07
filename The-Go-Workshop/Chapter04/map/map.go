package main

import (
	"fmt"
)

// A function that returns a map with string keys and string values
func getUsers() map[string]string {
	// Initialize a map with string keys and string values and then initialize with some elements
	users := map[string]string{"305": "Kofi", "204": "Asamoah", "530": "Akwasi"}
	// Add a new element to the map
	users["100"] = "Abena"
	return users
}
func main() {
	fmt.Println("Users:", getUsers())
}
