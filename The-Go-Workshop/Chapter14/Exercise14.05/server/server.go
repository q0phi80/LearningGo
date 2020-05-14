package main

import (
	"log"
	"net/http"
	"time"
)

/*
This creates a very basic web server that receives a request,
checks the authorization header is correct, waits 10 seconds, then sends back data.
*/
type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized"))
		return
	}
	time.Sleep(10 * time.Second)
	msg := "Hello Client!"
	w.Write([]byte(msg))
}
func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
