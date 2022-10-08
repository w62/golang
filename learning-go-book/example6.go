package main

import (
	"fmt"
)

const ( // start with 0 C preprocessor? 🤔
	a = iota
	b
	c
)

const ( //start with 1
	_ = iota
	a2 
	b2
	c2
)

const ( //start with 3
	_ = iota
	_ = iota
	_ = iota
	a3 
	b3
	c3
)

func main () {

//	var people [3]person <- failed to compile

	type person struct {
		name string
		age int8
	}

// compile ok
	var people [3]person
	var alice person

	fmt.Println(people)
	fmt.Printf("value: %v type: %T\n\n",people, people)
	fmt.Printf("value: %v type: %T", alice, alice)

	var x = new(int)

	a := [...]int{1, 2, 3}
	b := a
	c := &a
	b[1] = 5

	d := []int{1, 2, 3}
	e := &d


	fmt.Printf("value: %v type: %T\n\n", a, a)
	fmt.Printf("value: %v type: %T\n\n", d, d)
	fmt.Printf("value: %v type: %T\n\n", e, e)
	fmt.Printf("value: %v type: %T\n\n", *e, *e)
	fmt.Printf("value: %v type: %T\n\n", *e, *e)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x == nil)
	fmt.Println(*x)
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println()
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

	fmt.Println()
	fmt.Println(a3)
	fmt.Println(b3)
	fmt.Println(c3)


}
