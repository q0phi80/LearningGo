# The go build Command
### Syntax: go build main.go OR go build -o hello main.go
- This command compiles the application, including any packages and their dependencies, without installing the results.
- This creates a binary file to be executed.
- To reduce the file size of the produced binary, you can add additional flags during build
-- *e.g. go -o hello -ldflags "-w -s" main.go*

# Cross-Compiling
The build command allows for cross-compiling a go program for multiple Operating Systems and architectures.
- Reference: https://golang.org/doc/install/source#environment/
You need to set a constraint to cross-compile a go program.
The constraints include GOOS (the Operating System) and GOARCH (the architecture)
Constraints can be introduced via command line, code comments of a file suffix naming convention
Example cross-compile for Windows OS on a 64-architecture:
- *GOOS="windows" GOARCH="amd64" go build -o akwaaba maing.go*