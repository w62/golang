package main

// based on the tutorial here:
// https://go.dev/doc/tutorial/generics

/*
Test two more things:
- Define an interface recursively with another interface
- try to mix numbers and strings
*/

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

// Define an interface recursively with another interface
type AlphaNumeric interface {
	Number | rune | string
}


// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}

	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}

	strings := map[string]string{
		"first": "thirty five point nine eight ",
		"second": "twenty six point nine nine ",
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type paramters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

		fmt.Printf("Generic Sums with Constraints and more types:\nint: %v, \nfloat: %v and \nstring: %v\n",
		SumAlphaNumeric(ints),
		SumAlphaNumeric(floats),
		SumAlphaNumeric(strings))
}

//SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

// SumNumbers sums the values of map m. It supports both integers
// and flaots as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// try to mix numbers and strings
// SumAlphaNumeric sums the numeric values but concats the string
func SumAlphaNumeric[K comparable, V AlphaNumeric](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
