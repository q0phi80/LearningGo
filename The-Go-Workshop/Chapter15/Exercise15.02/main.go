package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello, world!</h1>"
	w.Write([]byte(msg))
}
func main() {
	// Use handle to route "/chapter1" through a handlefunc() function
	// We associate the path, /chapter1, with a function that returns a specific message (e.g. Chapter 1)
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "<h1>Chapter 1</h1>"
		w.Write([]byte(msg))
	})
	http.Handle("/", hello{})
	// Set the server to listen on a port number
	log.Fatal(http.ListenAndServe(":8080", nil))
}
