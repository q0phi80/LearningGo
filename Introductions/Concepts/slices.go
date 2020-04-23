package main

import (
  "fmt"
)

// Slices are similar to arrays in some other programming languages.
func main() {
    s := []int{1, 2, 3, 4} //Use square brackets to designate a slice and a type(e.g int) to indicate what the slice is going to contain.
    s = append(s, 5, 6, 7) //Use append to append new items to the slice.
    fmt.Println(s)
}
