package main

import (
  "fmt"
)
// Map is used for key value pairs.
func main() {
    m := map[string]string{} //Creating a map by putting the key value type in square brackets[] and the value type in curly braces{}
    m["first"] = "Kofi" // We can index the map and provide a key and assign a value
    m["last"]  = "Asamoah"
    fmt.Println(m)
}
