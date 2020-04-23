package main

import (
  "fmt"
)

// Every type in Go implements the empty interface.
func main() {
    var x interface{} // We can create a variable with an empty interface.
    x = "Hello, World!" // We can store any type in the interface.
    //x = 100 // We can assign any time such as integer.
    s, ok := x.(string) // If we are not sure of the type variable, we can supply a second variable ok and check our type variable is correct or not.
    //s, ok := x.(int)
    if !ok {
        panic ("NO!")
    }
    fmt.Println(s)
}
