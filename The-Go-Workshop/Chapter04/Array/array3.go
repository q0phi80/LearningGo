package main

import (
	"fmt"
)

// A function the defines an array with words
func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	// Create a message by joining the words in specific order and return it.
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}
func main() {
	fmt.Print(message())
}
