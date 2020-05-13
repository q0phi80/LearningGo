package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	/*
		Create database connection
	*/
	db, err := sql.Open("postgres", "user=postgres password=Fluffy123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[+] The connection to the DB was successfully initialized!")
	}
	// Periodically check whether the database still reachable
	connectivity := db.Ping()
	if connectivity != nil {
		panic(err)
	} else {
		fmt.Println("[+] Good to go!")
	}
	/*
	   Insert data into the database table
	*/
	insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}
	_, err = insert.Exec(2, "second")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[+] The value was successfully inserted!")
	}
	// Close the database once we are done running the application
	db.Close() // OR defer db.Close()
}
