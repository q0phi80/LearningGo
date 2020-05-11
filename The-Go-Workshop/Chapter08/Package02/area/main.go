package main

import (
	"LearningGo/The-Go-Workshop/Chapter08/Package02/shape"
)

// Set the fields for triangle, rectangle and square
func main() {
	t := shape.Triangle{Base: 15.5, Height: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t, r, s)
}
