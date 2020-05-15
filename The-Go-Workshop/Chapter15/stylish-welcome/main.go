package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	// Code to handle the static files
	http.Handle(
		"/statics/",
		http.StripPrefix(
			"/statics/",
			http.FileServer(http.Dir("./public")),
		),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
