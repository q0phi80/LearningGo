package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	// Define a map, initialize with data and return the map
	return map[string]string{"305": "Kwame", "204": "Yaw", "650": "Kwabena", "100": "Adwoa"}
}

// A function to accept a string as input...Return a string and Boolean
func getUser(id string) (string, bool) {
	users := getUsers()       // Get a copy of the users map from our earlier function
	user, exists := users[id] // Get a value from the users map using the passed ID as a key. Capture both the value and the exists value
	return user, exists       // Return both values
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("[!] User ID not passed.")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser(userID)
	// If the key is not found, print a message and then print all the users using the range loop
	if !exists {
		fmt.Printf("[!] Passed user ID (%v) not found.\nUsers:\n", userID)
		for key, value := range getUsers() {
			fmt.Println("	ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Println("Name:", name)
}
