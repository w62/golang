package main
// https://go.dev/doc/tutorial/create-module
// https://go.dev/doc/tutorial/call-module-code

// https://go.dev/doc/tutorial/handle-errors
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and 
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
