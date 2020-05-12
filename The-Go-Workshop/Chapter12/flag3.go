package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	i := flag.Int("age", -1, "Your age.")
	n := flag.String("name", "", "Your first name.")
	b := flag.Bool("married", false, "Are you married?")
	flag.Parse()
	if *n == "" {
		fmt.Println("Name is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("Name: ", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *b)
}
