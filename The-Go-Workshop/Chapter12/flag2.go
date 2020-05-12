package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.Int("age", -1, "Your age.")
	n := flag.String("name", "", "Your first name.")
	b := flag.Bool("married", false, "Are you married?")
	flag.Parse()
	fmt.Println("Name: ", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *b)
}
