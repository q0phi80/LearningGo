package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Custom error for when a file is not found
var (
	ErrWorkingFileNotFound = errors.New("[!] The working file is not found.")
)

// A function to perform a backup.
func createBackup(working, backup string) error {
	// Check to see if our working file exists,
	// before backing up
	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		}
		return err
	}
	// Open the working file and store the os.File returned by the function to the workFile variable
	workFile, err := os.Open(working)
	if err != nil {
		return err
	}
	// Read the content of the workFile
	content, err := ioutil.ReadAll(workFile)
	if err != nil {
		return err
	}
	// Write into a backup file
	// If the backup file does not exist, it will create the file
	// If the backup file does exist, it will overwrite the contents in the file
	err = ioutil.WriteFile(backup, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return nil // We return nil to indicate that we have not encountered any errors
}

// A function that will append data to the working file
// This function will accept the location of the working file
func addNotes(workingFile, notes string) error {
	notes += "\n"
	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // We will create the workingFile if it doesn't exist
	if err != nil {
		return err
	}
	defer f.Close()                                   // We need to close the file after the function closes.
	if _, err := f.Write([]byte(notes)); err != nil { // Write the contents of the note to the workingFile variable
		return err
	}
	return nil
}
func main() {
	backupFile := "backupFile.txt"
	workingFile := "note.txt"
	data := "note"
	// Call the createBackup() function to back up workingFile
	err := createBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 1; i <= 10; i++ {
		note := data + " " + strconv.Itoa(i)
		err := addNotes(workingFile, note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
