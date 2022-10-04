package main
import (
	"fmt"
)

func main () {
//	var x []int
x := make ([]int, 0, 5)
	fmt.Println(x, len(x), cap(x))
	x = append (x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append (x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append (x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append (x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append (x, 50)
	fmt.Println(x, len(x), cap(x))
}
