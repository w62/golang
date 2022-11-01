package isogram

import (
"strings"
)

func IsIsogram(word string) bool {

	var key byte

	word = strings.ToUpper(word)

	for i := 0; i < len(word); i++ {

		key = word[i]

		for j := i+1; j < len(word); j ++ {
		if word [i] != ' ' &&  word [i] != '-' {
			if key == word[j] {
				
				return false 
			}
		}
	}
}
	return true

	panic("Please implement the IsIsogram function")
}
