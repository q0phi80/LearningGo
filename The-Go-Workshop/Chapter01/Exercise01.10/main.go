package main

import (
    "fmt"
)
func main() {
    count := 5 // Create a variable with an initial value
    count += 5 // Add to it and then assign the result back to itself.
    fmt.Println(count)
    count++ // Increment the value by 1
    fmt.Println(count)
    count-- // Decrement the value by 1
    fmt.Println(count)
    count -= 5 // Subtract and assign the result back to itself.
    fmt.Println(count)

    name := "Kofi"
    name += " Asamoah" // Append another string to the end of the initial string.
    fmt.Println("Hello,", name)
}
