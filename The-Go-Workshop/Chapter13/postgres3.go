package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var prop string
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
	   Creating tables
	*/
	TableCreate := `
CREATE TABLE Number
(
	number integer NOT NULL,
	property text COLLATE pg_catalog."default" NOT NULL
)
WITH (
	OIDS = FALSE
)
TABLESPACE pg_default;
ALTER TABLE Number OWNER to postgres;`
	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[+] The table was successfully created!")
	}
	// Insert the numbers
	insert, insertErr := db.Prepare("INSERT INTO Number VALUES($1, $2)")
	if err != nil {
		panic(insertErr)
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "Even"
		} else {
			fmt.Println("The number:", i, " is:", prop)
		}
	}
	insert.Close()
	fmt.Println("The numbers are ready.")
	// Close the database once we are done running the application
	db.Close() // OR defer db.Close()
}
