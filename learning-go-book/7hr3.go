package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42 // 𨳒，咁都得
	// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=4:15:11
	fmt.Println(ms.foo) // 𨳒，咁都得
}

type myStruct struct {
	foo int
}
