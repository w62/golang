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
}

