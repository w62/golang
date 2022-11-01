package scrabble

import (
	"fmt"
)

func Score(word string) int {
	val := 0
	for i := 0; i < len(word); i ++ {
		val += lScore(word[i])
	}
	return val
	fmt.Println(word)
	panic("Please implement the Score function")
}

// function lScore returns the corresponding score of individual letter.
func lScore(b byte) int {

	val := 0

	switch {
	case b == 'a' || b == 'A':
		val = 1
	case b == 'e' || b == 'E':
		val = 1
	case b == 'i' || b == 'I':
		val = 1
	case b == 'o' || b == 'O':
		val = 1
	case b == 'u' || b == 'U':
		val = 1
	case b == 'l' || b == 'L':
		val = 1
	case b == 'n' || b == 'N':
		val = 1
	case b == 'r' || b == 'R':
		val = 1
	case b == 's' || b == 'T':
		val = 1
	case b == 't' || b == 'T':
		val = 1

	case b == 'd' || b == 'D':
		val = 2
	case b == 'g' || b == 'G':
		val = 2


	case b == 'b' || b == 'B':
		val = 3
	case b == 'c' || b == 'C':
		val = 3
	case b == 'm' || b == 'M':
		val = 3
	case b == 'p' || b == 'P':
		val = 3

	case b == 'f' || b == 'F':
		val = 4
	case b == 'h' || b == 'H':
		val = 4
	case b == 'v' || b == 'V':
		val = 4
	case b == 'w' || b == 'W':
		val = 4
	case b == 'y' || b == 'Y':
		val = 4

	case b == 'k' || b == 'K':
		val = 5

	case b == 'j' || b == 'J':
		val = 8
	case b == 'x' || b == 'X':
		val = 8

	case b == 'q' || b == 'Q':
		val = 10
	case b == 'z' || b == 'Z':
		val = 10

	default:
	}
	return val
}
