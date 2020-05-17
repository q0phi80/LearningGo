package exercise1

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

type UserDetails struct {
	Id         string
	CardNumber string
	Address    string
}

var testData = []*UserDetails{
	{
		Id:         "1",
		CardNumber: "1234",
		Address:    "123 Main Street",
	},
	{
		Id:         "2",
		CardNumber: "5678",
		Address:    "321 Blah Stret",
	},
}

// Function UpdatePhone() takes a sql.DB object and some user information such as an ID and a phone number as a string
// The UpdatePhone() function inserts a user ID and a phone number into the table by concatenating the data from the input parameters
// This function is vulnerable to SQLi
func UpdatePhone(db *sql.DB, Id string, phone string) error {
	var builder strings.Builder
	builder.WriteString("UPDATE USER_DETAILS SET PHONE=")
	builder.WriteString(phone)
	builder.WriteString(" WHERE USER_ID=")
	builder.WriteString(Id)
	fmt.Printf("Running query: %s\n", builder.String())
	_, err := db.Exec(builder.String())
	if err != nil {
		return err
	}
	return nil
}

//  Instead of concatenating inputs to form the query, use placeholders for the parameters to pass into the query
func UpdatePhoneSecure(db *sql.DB, Id string, phone string) error {
	stmt, err := db.Prepare(`UPDATE USER_DETAILS SET PHONE=? WHERE USER_ID=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(phone, Id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("[!] No row affected")
	}
	if rows > 1 {
		return fmt.Errorf("more than one row affected")
	}
	return nil
}

// A function to set up the database and load some test data
func initializeDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS USER_DETAILS (USER_ID TEXT, PHONE TEXT, ADDRESS TEXT)`)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(`INSERT INTO USER_DETAILS (USER_ID, PHONE, ADDRESS) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	for _, user := range testData {
		_, err := stmt.Exec(user.Id, user.CardNumber, user.Address)
		if err != nil {
			return err
		}
	}
	return nil
}

// Clear the database
func tearDownDB(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE USER_DETAILS")
	if err != nil {
		return err
	}
	return nil
}

// Set up the database connection
func getConnection() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "test.DB")
	if err != nil {
		return nil, fmt.Errorf("[!] Could not open db connection %v", err)
	}
	return conn, nil
}

// Execute the setup of the test data and then run the test.
func TestMain(m *testing.M) {
	var err error
	db, err = getConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = initializeDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tearDownDB(db)
	if m.Run() != 0 {
		fmt.Println("[!] Error running test.")
		os.Exit(1)
	}
}

// Run a test against the UpdatePhoneSecure() function
func TestUpdatePhoneSecure(t *testing.T) {
	var tests = []struct {
		ID    string
		Phone string
		err   string
	}{
		{
			ID:    "1",
			Phone: "1234",
			err:   "",
		},
		{
			ID:    "'2' OR USER_ID=='1'",
			Phone: "5678",
			err:   "[!] No row affected",
		},
	}
	for _, test := range tests {
		err := UpdatePhoneSecure(db, test.ID, test.Phone)
		if err != nil {
			assert.EqualError(t, err, test.err)
		}
	}
}
