package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("Born on a weekend.")
	default:
		fmt.Println("Born some other day.")
	}
}
