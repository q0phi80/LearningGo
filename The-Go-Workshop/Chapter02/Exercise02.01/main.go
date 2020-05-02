package main

import (
	"fmt"
)

func main() {
	input := 9
	if input%2 == 0 {
		fmt.Println(input, "is even.")
	}
	if input%2 == 1 {
		fmt.Println(input, "is odd")
	}
}
