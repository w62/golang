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

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)

	d, err := divide(5.0, 1.0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)

	f:= func() {
			fmt.Println("hello go")
	}

	f()

}


func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func dump (a, b float64) (float64, float64, error) {
	return a, b, nil
}

type myStruct struct {
	foo int
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return &result

	// Go will promote stack to heap when returning a pointer
	// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=4:35:30
	
}
