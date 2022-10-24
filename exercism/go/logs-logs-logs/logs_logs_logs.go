package logs

import (
//	"fmt"
//	"unicode/utf8"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	results := ""
	//r,_ := utf8.DecodeRuneInString(log)

	if strings.IndexRune(log, '\u2757') ==0 {
		results = "recommendation"
	} else if strings.IndexRune(log, '\U0001f50d') ==0 {
		results = "search"
	} else if strings.IndexRune(log, '\u2600') ==0 {
		results = "weather"
	} else if strings.IndexRune(log, '\u2600') >=0 {
		results = "weather"
	} else if strings.IndexRune(log, '\U0001f50d') >=0 {
		results = "search"
	} else if strings.IndexRune(log, '\u2757') >=0 {
		results = "recommendation"
	} else {
		results = "default"
	}


	return results
	panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
