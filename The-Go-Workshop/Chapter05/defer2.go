package main

import (
	"fmt"
)

func main() {
	defer func() { // Defer before anonymous function
		fmt.Println("Now I am done")
	}()
	fmt.Println("Main:	Start")
	fmt.Println("Main:	End")
}
