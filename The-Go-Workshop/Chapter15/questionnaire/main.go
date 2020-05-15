package main

import (
	"html/template"
	"log"
	"net/http"
)

// A struct to hold all the attributes for the template
type Visitor struct {
	Name    string
	Surname string
	Age     string
}
type Hello struct {
	tpl *template.Template
}

// Create a handler function for the handler
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vst := Visitor{} // Create a new empty visitor
	// Check if the request is a post
	if r.Method == http.MethodPost {
		// Parse the form
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		// If the form parsed correctly, assign the parameters to the visitor's attribute
		vst.Name = r.Form.Get("name")
		vst.Surname = r.Form.Get("surname")
		vst.Age = r.Form.Get("age")
	}
	h.tpl.Execute(w, vst)
}

// Create a handler
func NewHello(tplPath string) (*Hello, error) {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}
	return &Hello{tmpl}, nil
}
func main() {
	hello, err := NewHello("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", hello)
	http.HandleFunc("/form", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./form.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
