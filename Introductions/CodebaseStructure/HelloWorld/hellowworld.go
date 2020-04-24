package helloworld

import "golang.org/x/sys/unix"

// The function is capitalized because it is going to be a public function.
// GetHellowWorld will get the hello world string.
func GetHellowWorld() string { // Going to return a string.
    return "Hello, World!" // The string the function is going to return.

}

// GetUserID gets the ID of the current user using native Unix OS utility
func GetUserID() int {
    return unix.Getuid()
}
