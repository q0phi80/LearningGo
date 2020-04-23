package main

import (
  "fmt"
)
func main() {
  // Iterate over a list of values
  for _, i := range []int{1,2,3,4} {// Created a list of integer values.Integer arrays. We pass the integer list to the range keyword
      fmt.Println(i)
  }
}
