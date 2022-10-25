package main

import (
	"fmt"
)

func main () {
	var i int

	i = 10

//	v, ok := i.(int)
//	if ok {
		fmt.Printf("i is type %T, value %d\n", i, i)
//	} else {
//		fmt.Printf("unknown type")
//	}
}
