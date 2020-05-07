package main

import (
	"fmt"
)

// A function that fills an empty array with numbers from 1 to 10.
func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

// A function that multiplies the number from an array by itself and then sets the result back to the array
func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

// The main() defines our empty array, fill it, modify it and then print its contents to the console.
func main() {
	var arr [10]int //An empty array
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}
