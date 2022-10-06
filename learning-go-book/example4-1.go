package main
import (
	"fmt"
	"math/rand"
)

func main () {
	if n:= rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}


	if name := "mary"; name == "mary" {
		fmt.Println("it's mary!")
	} else if name == "peter" {
		fmt.Println("it's peter!") 
	} else {
		fmt.Println("who are you?", name)
	}

	for i := 0; i < 10; i ++ {
		fmt.Println(i)
	}

	// The scope of i is the for loop. It goes away after the block


// That is why we need another i here.
 	var i int = 0

	for i < 10 {
		fmt.Println("2nd for:", i)
		i ++
	}

	i = 0

	for {
		fmt.Println ("Hello")

		i++
		if !(i < 20) {
			fmt.Println("88 after", i, "times.")
			break
		}
	}

	for j := 0; j < 20; j ++ {
		if j % 3 == 0 {
			continue
		}
		fmt.Println("here: ", j)
		fmt.Println("::::: ")

	}

	for k := 1; k <= 100; k ++ {
		if k%3 == 0 && k%5 == 0 {
			fmt.Println("FizzBuzz", k)
			continue
		}

		if k%3 == 0 {
			fmt.Println("Fizz", k)
			continue
		}

		if k%5 == 0 {
			fmt.Println("Buzz", k)
			continue
		}

		fmt.Println(k)
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}

	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true,
	"Raul": true, "Wilma": true}

	for k := range uniqueNames {
		fmt.Println(k)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	samples := []string{"hello", "apple_🍀!"}
	for _, sample := range samples {
		for i, r := range sample {
		fmt.Println(i, r, string(r))
		}
	fmt.Println()
	}

	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals2 {
		v *= 2
	}

	fmt.Println(evenVals)

	samples = []string{"hello", "apple_🍀"}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	words = []string{"hi", "salutations", "hello"}

	for _, word := range words {
		// can include a short variable declaration
		switch wordLen := len(word) ; {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}


	var a int = 5


// this is no good	
	switch {
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	case a == 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("otherwise, a is ", a)
	}

// this is better
	switch a {
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("otherwise, a is ", a)
	}

	a = rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}

	fmt.Println("do something when the loop completes normally")
	done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

		

// The closing bracket of the main () function
}
