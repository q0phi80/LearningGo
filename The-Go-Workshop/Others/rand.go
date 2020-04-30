package main

import (
  "math/rand"
)

func main() {
    //var seed = 123456789 // This will throw an error because rand.Seed requires a variable of the int64 type.
    var seed int64 = 123456789
    rand.Seed(seed)
}
