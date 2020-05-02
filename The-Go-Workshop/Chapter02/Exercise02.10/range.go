package main

import (
	"fmt"
)

func main() {
	// Define a map with a string key and a string value of strings
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}
	// Use range to get the key and value for an array element and assign them to variables
	for key, value := range config {
		// Print out the key and value variables
		fmt.Println(key, "=", value)
	}
}
