package main

import (
  "fmt"
)
// We can also create map by using the make and slice keywords.
func main() {
    m := make(map[string]string) // Create an empty map.
    // Index and assign
    m["First"] = "Kofi"
    m["Last"]  = "Asamoah"
    fmt.Println(m["Last"]) // We can also index to retrieve a particular value. In this case, we want to retrieve just the last name.
}
