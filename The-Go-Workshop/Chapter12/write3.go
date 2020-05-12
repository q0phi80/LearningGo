package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Stat("junk.txt") // We call os.Stat() on the junk.txt file to check if it exists. os.Stat() method returns a FileInfo type if the file exists, otherwise it will return nil.
	if err != nil {
		if os.IsNotExist((err)) { // The os.Stat() method can return multiple errors so we must check to confirm that the error generated is due to the absence of the file.
			fmt.Println("junk.txt: File does not exist!")
			fmt.Println(file) // This will be nil since the file doesn't exist.
		}
	}
	fmt.Println()
	file, err = os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist((err)) {
			fmt.Println("test.txt: File does not exist!")
		}
	}
	fmt.Printf("file name: %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize: %d\n", file.Name(), file.IsDir(), file.ModTime(), file.Mode(), file.Size())
}
