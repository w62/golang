package diffsquares
import (
//	"fmt"
)

func SquareOfSum(n int) int {
	j := 0

	for i := 1; i <=n; i ++ {
		j += i
	}

	return j * j
	panic("Please implement the SquareOfSum function")
}

func SumOfSquares(n int) int {

	j := 0

	for i := 1; i <=n; i ++ {
		j = j + i * i
	}
	return j
	panic("Please implement the SumOfSquares function")
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
	panic("Please implement the Difference function")
}
