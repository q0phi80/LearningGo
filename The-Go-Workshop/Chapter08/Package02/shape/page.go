package shape

import "fmt"

// A Shape{} interface with two method sets called Area() float64 and Name() string
type Shape interface {
	Area() float64
	Name() string
}

// Create struct types
// Each of the structs will satisfy the Shape{} interface
type Triangle struct { // Capitalize triangle's T to make exportable.
	Base   float64
	Height float64
}
type Rectangle struct {
	Length float64
	Width  float64
}
type Square struct {
	Side float64
}

// Create the Area() and Name() methods for the triangle structure type
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2 // The area of a triangle is base*height/2
}
func (t Triangle) Name() string {
	return "Triangle" // The Name() method returns the name of the shape.
}

// Create the Area() and Name() methods for the rectangle struct type
func (r Rectangle) Area() float64 {
	return r.Length * r.Width // The area of rectangle is length*width
}
func (r Rectangle) Name() string {
	return "Rectangle" // The Name() method returns the name of the shape
}

// Create the Area() and Name() methods for the square struct type
func (s Square) Area() float64 {
	return s.Side * s.Side // The area of a square is side*side
}
func (s Square) Name() string {
	return "Square"
}

// A function that accepts the Shape interface as a variadic parameter
// The function will iterate over the Shape type and will execute each of its Name() and Area() methods
func PrintShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}
