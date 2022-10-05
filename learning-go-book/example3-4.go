package main

import (
	"fmt"
)

func main () {
	var s string = "Hello 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	var nilMap map[string]int
	fmt.Println(bs)
	fmt.Println(rs)
	fmt.Println(nilMap)

	teams := map[string][]string {
		"Orcas": []string {"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams)

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])


	m := map[string]int {
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	delete(m, "hello")

	v, ok = m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)


	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1 ,2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}


	type person struct {
		name string
		age int
		pet string
	}

	var fred person
	bob := person {}
	julia := person {
		"Julia",
		40,
		"cat",
	}

	beth := person{
		age: 30,
		name: "Beth",
	}


	fmt.Println(fred)
	fmt.Println(bob)
	fmt.Println(julia)
	fmt.Println(beth)

	bob.name = "Bob"
	fmt.Println(beth.name)
	fmt.Println(bob)


	var ninja struct {
		name string
		age int
		pet string
	}

	ninja.name = "unknown"
	ninja.age = 100
	ninja.pet = "unknown"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(ninja)
	fmt.Println(pet)

	type firstPerson struct {
		name string
		age int
	}

	type secondPerson struct {
		name string
		age int
	}

	type thirdPerson struct {
		age int
		name string
	}

	type forthPerson struct {
		firstName string
		age int
	}

	type ffthPerson struct {
		name string
		age int
		favoriteColor string
	}

	f := firstPerson {
		name: "Bob",
		age: 50,
	}

	var g struct {
		name string
		age int
	}

	g = f

	// compiles -- can use = and == between identical
	// named and anonymous structs

	fmt.Println(f == g)
}
