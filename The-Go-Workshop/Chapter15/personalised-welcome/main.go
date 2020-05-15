package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// A function that can be used as a handling function for an HTTP path
func Hello(w http.ResponseWriter, r *http.Request) {
	// Save the query to a variable using the Query method URL from the request
	v1 := r.URL.Query()
	name, ok := v1["name"] // The second variable, ok, is a Boolean that tells us whether the "name" key exists
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s", strings.Join(name, ","))))
}
func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
