package main

import (
	"fmt"
)

func main () {
	var s string = "hello"
	var s2 []string
	fmt.Printf("s=%s type=%T type s[0]=\n", s, s, s[0])
	fmt.Println("zero value of s2=",s2)
}
