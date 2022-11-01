package main

import (
	"fmt"
)

func main () {
	var s string = "hello"
	fmt.Printf("s=%s type=%T type s[0]=\n", s, s, s[0])
}
