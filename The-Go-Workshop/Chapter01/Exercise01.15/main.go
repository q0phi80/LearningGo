package main

import (
    "fmt"
)
// Create a function that accepts an int as an argument
func add5Value(count int) {
    count += 5 // Add 5 to the passed number.
    fmt.Println("add5Value    :", count) // Print the updated number to the console.
}
// Create a function that takes an int pointer
func add5Point(count *int) {
    *count += 5 // Dereference (place * in front of the variable) the value and add 5 to it.
    fmt.Println("add5Point    :", *count) // Print the updated value of the count and dereference it.
}
func main() {
    var count int // Declare an int variable
    add5Value(count) // Call the first function with the variable
    fmt.Println("add5Value post:", count) // Print the current value of the variable.
    add5Point(&count) // Since we are a calling a pointer, u add & before the variable when calling the second function.
    fmt.Println("add5Point post:", count)
}
