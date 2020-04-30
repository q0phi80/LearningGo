package main

import (
	"fmt"
)

var foo string = "bar" // Declaration at the package level

func main() {
	var baz string = "qux" // Declaration at the function level
	fmt.Println(foo, baz)
}
