package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `firstName, lastName, age
	Asamoah, Gyan, 55
	Kwadwo, Asamoah, 45
	Akwasi, Amponsah, 2`

	r := csv.NewReader(strings.NewReader(in)) // Create a reader type and return it.
	for {
		record, err := r.Read() // We read each record one at a time.
		if err == io.EOF {      // Check if we have reached the end of the file.
			break // Break out of the loop
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
