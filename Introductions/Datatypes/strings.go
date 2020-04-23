package main

import (
  "fmt"
  "strings"
)

func main () {
    s := "I'm a string"
    fmt.Printf("Ends with string? %v\n", strings.HasSuffix(s, "string")) // Using a function called HasSuffix from the strings package to return a boolean value after checking if our string s ends with the word "string"
}
