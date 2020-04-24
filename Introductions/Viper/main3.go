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
    viper.AddConfigPath(".")
    viper.SetConfigName("creds")
    viper.ReadInConfig()
    /*if err := viper.ReadInConfig(); err != nil {
        log.Fatalln("[!] Unable to read config.")
    }*/
    if username != "" {
        viper.Set("credentials.username", username)
    }
    if password != "" {
        viper.Set("credentials.password", password)
    }
    if !viper.IsSet("credentials.username") || !viper.IsSet("credentials.password") {
        log.Fatalln("[!] No credentials supplied.")
    }
    useCredentials()
}
