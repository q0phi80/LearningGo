package main

import (
  "fmt"
)
// Pointers can also be used for other things such as integers.
func main() {
  i := new(int)
  fmt.Println(i)
  fmt.Println(*i)
}
