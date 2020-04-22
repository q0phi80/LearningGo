package main

import (
  "fmt"
)
func main() {
  a := 10

  switch a == 10 { //Checking if a is equal to 10
  case true: // if a is equal to 10, then print the message below.
      fmt.Println("Yes!")
  case false: // if a is not equal to 10, then print the message below.
      fmt.Println("No!")    
  }
}
