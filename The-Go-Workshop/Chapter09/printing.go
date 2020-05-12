package main

import (
	"fmt"
)

func main() {
	fname := "q0phi80"
	gpa := 3.75
	hasJob := true
	age := 24
	hourlyWage := 45.53
	fmt.Printf("%s has gpa %.2f.\n", fname, gpa)
	fmt.Printf("He has a job equals %t.\n", hasJob)
	fmt.Printf("Hes is %d earning %v per hour.\n", age, hourlyWage)
}
