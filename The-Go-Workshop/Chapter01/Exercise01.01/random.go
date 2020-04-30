package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())
    // Generate a random number between 0 and 4 then add 1 to get a number between 1 and 5.
    r := rand.Intn(5) + 1
    // Use the string repeator to create a string with the number of stars we need.
    stars := strings.Repeat("*", r)
    //Print the string
    fmt.Println(stars)
}
