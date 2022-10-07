package main

import (
	"fmt"
	"errors"
	"os"
	"strconv"
	"time"
	"sort"
	"io"
	"log"
)

type person struct {
	age int
	name string
}

func modMap(m map [int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m,1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s,10)
}

func main () {

	m := map[int]string{
		1: "first",
		2: "second",
	}

	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)


	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5} ...))

	 result, remainder, err := divAndRemainder(5 ,2)

	 if err != nil {
		 fmt.Println(err)
		 os.Exit(1)
	 }
	 fmt.Println(result, remainder ) 

	 type opFuncType func (int, int) int

	var opMap = map[string] opFuncType {
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of any anonymous function")
		}(i)

		func () {
			dt := time.Now()
			time.Sleep(time.Second)
			fmt.Println("yet another anonymous function. Time is:", dt.Format("01-02-2006 15:04:05:06"))
		}()
	}

	type Person struct {
		FirstName string
		LastName string
		Age int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	var base0 int = 2
	var base1 int = 3
	var factor int = 5

	var twoBase func(int) (int) = makeMult(base0)
	var threeBase func(int) (int) = makeMult(base1)

	fmt.Println(makeMult(base0)(factor))
	fmt.Println(makeMult(base1)(factor))

	fmt.Println(twoBase(factor))
	fmt.Println(threeBase(factor))

	for i := 0; i < 6; i++ {
		fmt.Println(i, twoBase(i), threeBase(i))
	}

<<<<<<< HEAD

	p := person{}
	i := 2
	s1 := "hello"

	fmt.Println(i, s1, p)
	modifyFails(i, s1, p)
	fmt.Println(i, s1, p)
=======
	if len(os.Args) < 2{
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
>>>>>>> refs/remotes/origin/main

}
// main ends here 

<<<<<<< HEAD
func modifyFails (i int, s string, p person) {
	i = i *2
	s = "Goodbye"
	p.name = "Bob"
=======

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	return file, func () {
		file.Close()
		fmt.Println("deferred close.")
	}, nil
>>>>>>> refs/remotes/origin/main
}

func makeMult(base int) func (int) int {
	return func(factor int) int {
		return base * factor 
	}
}

func addTo(base int, vals ...int)[]int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out

}

func divAndRemainder(numerator int, denominator int) ( result int, remainder int, err error)  {
	if denominator == 0 {
		err =  errors.New("cannot divide by zero")
		// no need to assign 0 to result or remainder as default is 0
		return 
	}
	
 result, remainder = numerator / denominator, numerator % denominator
 return
}

func add (i int, j int) int {return i + j}
func sub (i int, j int) int {return i - j}
func mul (i int, j int) int {return i * j}
func div (i int, j int) int {return i / j}


