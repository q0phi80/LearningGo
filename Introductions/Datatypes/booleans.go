package main

import (
  "fmt"
)

func main() {
    a := true // := is an assignment operator for a new variable that doesn't exist. This also means we don't have to tell Go what type the variable 'a' is.
    b := false

    fmt.Printf("a = %v, b = %v, and a == b = %v\n", a, b, a == b)
}
