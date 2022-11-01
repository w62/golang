package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	val := ""

	if number%3 == 0 {
		val += "Pling"
	}

	if number%5 == 0 {
		val += "Plang"
	}
	if number%7 == 0 {
		val += "Plong"
	}

	if number%3 != 0 &&
		number%5 != 0 &&
		number%7 != 0 {
		val = strconv.Itoa(number)
	}

	return val
	panic("Please implement the Convert function")
}
