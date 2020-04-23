package main

import (
  "fmt"
)
// We can use an empty interface with a switch statement
func main() {
    var x interface{} // Empty interface
    x = "Hello, World!"
    switch x.(type) {
    case int:
        fmt.Println("Integer!")
    case string:
        fmt.Println("String!")     

    }
}
