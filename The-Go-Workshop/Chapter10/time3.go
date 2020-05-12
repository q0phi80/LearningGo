package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	AppName := "HTTPCHECKER"
	Action := "BASIC"
	Date := time.Now()
	LogFileName := AppName + "_" + Action + "_" + strconv.Itoa(Date.Year()) + "_" + Date.Month().String() + "_" + strconv.Itoa(Date.Day()) + ".log"
	fmt.Println("The name of the logfile is: ", LogFileName)
}
