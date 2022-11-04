// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b <= c || b+c <= a || c+a <= b {
		return NaT
	}

	if a == b && b == c && c == a {
		k = Equ
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	k = Sca

	return k
}
