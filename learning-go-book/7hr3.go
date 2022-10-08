package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42 // deferencing has lower priorty of the dot notation
	// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=4:14:31
	fmt.Println((*ms).foo)
}

type myStruct struct {
	foo int
}
