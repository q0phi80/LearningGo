package main

import (
	"fmt"
)

// Define an array with 10 elements
func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}
func main() {
	var arr [10]int      // Define an empty array
	arr = fillArray(arr) // Call the fillArray() function.
	fmt.Println(arr)     // Print the element
}
