package main

import (
	"fmt"
)


func main () {

	var 叉油雞飯 int8 = 2

	var str string = "ABCD"
	var x = []int{1, 5: 4, 6, 10:100, 15}
	var y = [...]int{1, 5: 4, 6, 10:100, 15}
    r_array := []rune(str)
    fmt.Printf("Array of str values for A, B, C and D: %U\n", str)
    fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)
	fmt.Println(叉油雞飯)
	fmt.Println(x[2])
	fmt.Println("length is: ", len(x))
	fmt.Println(y)
}
