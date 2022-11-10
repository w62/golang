// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var result []string
	if len(rhyme) <= 0 {
		return nil
	}
	if len(rhyme) == 1 {

		result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	} else {
		modern := true
		if rhyme[0] == "pin" {
			modern = true
		} else {
			modern = false
		}

		for len(rhyme) > 0 {
			if len(rhyme) >= 2 {
				result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[0], rhyme[1]))
				rhyme = rhyme[1:]
			} else if len(rhyme) == 1 {
				if modern {
					result = append(result, fmt.Sprintf("And all for the want of a pin."))
				} else {
					result = append(result, fmt.Sprintf("And all for the want of a nail."))
				}
				rhyme = rhyme[1:]

			}
		}
	}
	return result
	panic("Please implement the Proverb function")
}
