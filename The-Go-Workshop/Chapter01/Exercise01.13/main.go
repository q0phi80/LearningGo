package main

import (
    "fmt"
    "time"
)

func main() {
    // Declare a pointer using a var statement
    var count1 *int
    // Create a variable using new.
    count2 := new(int)
    // Create a temporary variable to hold a number.
    countTemp := 5
    // Use & to create a pointer from an existing variable.
    count3 := &countTemp
    // Creating a pointer from some types without a temporary variable
    t := &time.Time{}
    // Print each output
    fmt.Printf("Count1: %#v\n", count1) // This should produce a nil output since we didn't assign any value
    fmt.Printf("Count2: %#v\n", count2) // This will produce the memory address of the pointer
    fmt.Printf("Count3: %#v\n", *count3) // To get the value the pointer is associated with, we have to dereference the value using * in front of the variable name.
    fmt.Printf("Time : %#v\n", t) // This will produce a memory address.
}
