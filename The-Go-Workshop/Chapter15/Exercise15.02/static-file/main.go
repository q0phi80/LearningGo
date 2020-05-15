package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html") // Sends to ResponseWriter the content of the "index.html" file
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
