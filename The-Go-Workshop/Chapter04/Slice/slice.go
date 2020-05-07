package main

import (
	"fmt"
	"os"
)

// The function takes an int argument and returns a strinig slice
func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are required\n", minArgs)
		os.Exit(1)
	}
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

// The function loops over a slice and finds the longest string. When two words are of the same length, the first word wins.
func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

// The main() function calls the functions and check for errors.
func main() {
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error.")
		os.Exit(1)
	}
}
