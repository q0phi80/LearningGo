package main

import (
	"fmt"
	"os"
)

// Define a users map
var users = map[string]string{
	"305": "Kwame",
	"310": "Kwabena",
	"540": "Yaw",
	"631": "Ama",
	"051": "Abena",
}

// A function that deletes from the users map using a passed in string as the key
func deleteUser(id string) {
	delete(users, id)
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("[!] User ID not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	deleteUser(userID)
	fmt.Println("Users:", users)
}
