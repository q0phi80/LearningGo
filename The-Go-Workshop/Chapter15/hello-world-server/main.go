package main

import (
	"log"
	"net/http"
)

// Create a handler, the struct that handles the requests
type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello, World!</h1>"
	w.Write([]byte(msg))
}

// Main function will start the server
func main() {
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
