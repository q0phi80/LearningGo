package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 255; i++ { // Loop up to 255 times
		/*
			Display i as a decimal value with a width of 3 and right aligned.
			Display i as a base 2 value with a width of 8 and right aligned.
			Display i as a hex value with a width of 2 and right aligned.
		*/
		fmt.Printf("Decimal: %3.d Base Two: %8.d Hex: %2.x\n", i, i, i)
	}
}
