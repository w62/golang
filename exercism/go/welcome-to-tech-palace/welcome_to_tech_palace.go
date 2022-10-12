package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
//	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	s := "strings"
	s = strings.Repeat("*", numStarsPerLine)
	r := fmt.Sprintf("%s\n%s\n%s", s, welcomeMsg, s)
	return r
//	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	s := "strings"
	s = strings.ReplaceAll(oldMsg,"*", " ")
	s = strings.TrimSpace(s)
	return s
//	panic("Please implement the CleanupMessage() function")
}
