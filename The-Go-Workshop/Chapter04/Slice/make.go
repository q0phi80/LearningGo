package main

import (
	"fmt"
)

/*
Using the make function to create some slices and display their length and capacity
*/

// A function that returns three int slices
func genSlices() ([]int, []int, []int) {
	var s1 []int              // Define a slice using the var notation.
	s2 := make([]int, 10)     // Define a slice using make and set only the length
	s3 := make([]int, 10, 50) // Use make to define a slice that uses both the length and capacity.
	return s1, s2, s3         // Return the three slices
}
func main() {
	s1, s2, s3 := genSlices() // Call the function and capture the returned values
	// For each slice, print its length and capacity
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))

}
