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
	   Creating tables
	*/
	DBCreate := `
CREATE TABLE public.test
(
	id integer,
	name character varying COLLATE pg_catalog."default"
)
WITH (
	OIDS = FALSE
)
TABLESPACE pg_default;
ALTER TABLE public.test OWNER to postgres;`
	_, err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[+] The table was successfully created!")
	}
	// Close the database once we are done running the application
	db.Close() // OR defer db.Close()
}
