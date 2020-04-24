package main

import (
  "fmt"
  "flag"
  "log"
  "github.com/spf13/viper"
)
var username, password string

func useCredentials() {
    fmt.Printf("Your username is %s\n", viper.GetString("credentials.username"))
    fmt.Printf("Your password is %s\n", viper.GetString("credentials.password"))
}

func main() {
    flag.StringVar(&username, "user", "", "Username to use.")
    flag.StringVar(&password, "pass", "", "Password to use.")
    flag.Parse()
    if username == "" || password == "" {
        log.Fatalln("[!] Must pass credentials.")
    }
    viper.Set("credentials.username", username)
    viper.Set("credentials.password", password)
    useCredentials()
}
