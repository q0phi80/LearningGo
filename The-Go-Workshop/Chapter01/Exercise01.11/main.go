package main

import (
  "fmt"
)
func main() {
  // Define our visits variable
  visits := 15
  fmt.Println("First visit    :", visits == 1) // Use the equal signs to check if this is the first visit.
  fmt.Println("Return visit   :", visits != 1) // Use the not equal sign to check if this is a returning visitor.
  // Check if the visitor is a Silver member (btn 10 and 20 visits inclusively)
  fmt.Println("Silver member  :", visits >= 10 && visits < 21)
  fmt.Println("Gold member    :", visits > 20 && visits <= 30)
  fmt.Println("Platnum member :", visits > 30)
}
