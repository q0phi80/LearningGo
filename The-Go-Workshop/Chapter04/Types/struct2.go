package main

import (
	"fmt"
)

/*
Comparing structs to each other
*/

// Define a comparable struct
type point struct {
	x int
	y int
}

// A function that returns two Booleans
func compare() (bool, bool) {
	// Anonymous structs
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}
	point2 := struct {
		x int
		y int
	}{} // Initialize to zero
	point2.x = 10
	point2.y = 5

	// Using the named struct type created previously
	point3 := point{10, 10}

	// Compare them and return
	return point1 == point2, point1 == point3
}

// Call the struct function and print the results
func main() {
	a, b := compare()
	fmt.Println("point1 == point2:", a)
	fmt.Println("point1 == point3:", b)
}
