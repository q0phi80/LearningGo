Interface
An interface is a collection of methods that describe the behavior of the data type.
e.g.
type Speaker interface{
Speak(message string) string // A Speaker() method that accepts a string and returns a string.
Greet() string // A Greeting() method that returns a string
}
Interfaces in Go provide a way to specify the behavior of an object.
A behavior is defined by a set of methods.
A behavior describes what that type can do.
Behaviors of a type can be as follows:
Read()
Write()
Save()

It is idiomatic in Go to name the interface ending in er (e.g. type Speaker interface {})

An example of interface with method sets (multiple methods)
// https://golang.org/pkg/os/#FileInfo 
type FileInfo interface { 
        Name() string // base name of the file 
        Size() int64 // length in bytes for regular files; system-dependent for others 
        Mode() FileMode // file mode bits 
        ModTime() time.Time // modification time 
        IsDir() bool // abbreviation for Mode().IsDir() 
        Sys() interface{} // underlying data source (can return nil) 
} 

You can write interfaces for types in another package.