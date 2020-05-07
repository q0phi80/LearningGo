package main

import (
	"fmt"
	"os"
)

// A function to grab a user input
func getUserInput() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

// A function that accepts a slice of strings as a parameter and returns a slice of strings
func getLocals(extraLocals []string) []string {
	var locales []string                        // A slice of string variable
	locales = append(locales, "en_US", "fr_FR") // Add multiple strings to the slice using append
	locales = append(locales, extraLocals...)   // Add more data from the parameter
	return locales                              // Return the variable
}

// The main() function gets the user input, pass it to the getUserInput() function and print the result
func main() {
	locales := getLocals(getUserInput())
	fmt.Println("Locales to use:", locales)
}
