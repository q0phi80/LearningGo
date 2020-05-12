LOG
The log functions, Println(), Printf(), and Print(), perform the same functionality as their fmt counterparts with one exception. When the log functions execute, it gives additional details such as the date and time of the execution.

Using the log package, we can also log fatal errors. The Fatal(), Fatalf(), and Fatalln() functions are similar to Print(), Printf(), and Println(). The difference is after the log Fatal() functions are followed by an os.Exit(1) a system call.

The log package also has the following functions: Panic, Panicf, and Panicln. The difference between the Panic() functions and the Fatal function is that the Panic functions are recoverable. When using the Panic functions, you can use the defer() function, whereas when using the Fatal functions, you cannot.