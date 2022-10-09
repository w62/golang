package main
// https://go.dev/doc/tutorial/create-module
// https://go.dev/doc/tutorial/call-module-code
import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
