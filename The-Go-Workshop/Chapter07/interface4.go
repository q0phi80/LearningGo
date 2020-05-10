package main

import "fmt"

// A Shape{} interface with two method sets called Area() float64 and Name() string
type Shape interface {
	Area() float64
	Name() string
}

// Create struct types
// Each of the structs will satisfy the Shape{} interface
type triangle struct {
	base   float64
	height float64
}
type rectangle struct {
	length float64
	width  float64
}
type square struct {
	side float64
}

// Create the Area() and Name() methods for the triangle structure type
func (t triangle) Area() float64 {
	return (t.base * t.height) / 2 // The area of a triangle is base*height/2
}
func (t triangle) Name() string {
	return "triangle" // The Name() method returns the name of the shape.
}

// Create the Area() and Name() methods for the rectangle struct type
func (r rectangle) Area() float64 {
	return r.length * r.width // The area of rectangle is length*width
}
func (r rectangle) Name() string {
	return "rectangle" // The Name() method returns the name of the shape
}

// Create the Area() and Name() methods for the square struct type
func (s square) Area() float64 {
	return s.side * s.side // The area of a square is side*side
}
func (s square) Name() string {
	return "square"
}

// A function that accepts the Shape interface as a variadic parameter
// The function will iterate over the Shape type and will execute each of its Name() and Area() methods
func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}

// Set the fields for triangle, rectangle and square
func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	s := square{side: 10}
	printShapeDetails(t, r, s)
}
