package main

import (
  "fmt"
  "flag"
  "log"
)
var username, password string

func useCredentials(username, password string) {
    fmt.Printf("Your username is %s\n", username)
    fmt.Printf("Your password is %s\n", password)
}

func main() {
    flag.StringVar(&username, "user", "", "Username to use.")
    flag.StringVar(&password, "pass", "", "Password to use.")
    flag.Parse()
    if username == "" || password == "" {
        log.Fatalln("[!] Must pass credentials.")
    }
    useCredentials(username, password) // Function to pass in our username and password.
}
