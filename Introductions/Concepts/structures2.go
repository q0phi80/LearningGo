package main

import (
  "fmt"
)
// We can also create an Anonymous struct.
func main() {
    pt := struct {// Creating a point with x,y coordinates.
        X int
        Y int
      }{
          X: 10,
          Y: 9,
      }
      fmt.Println(pt)
}
