package main

import (
	"log"
	"net/http"
)

/*
A basic server that sends back JSON data.
*/
type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "{\"message\": \"hello world\"}"
	w.Write([]byte(msg))
}
func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
