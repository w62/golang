package grains

import (
	"errors"
	"math"
	"fmt"
)

func Square(number int) (uint64, error) {
	var val uint64 = 1
	var ival uint64 = 1
	var r uint64 = 2
	var n uint64 = uint64(number)

	if number < 1 || number > 64 {
		return 0, errors.New("invalid range")
	}

	val = ival * uint64(math.Pow(float64(r), float64(n - 1)))

	return val, nil
	panic("Please implement the Square function")
}

func Total() uint64 {
	var val uint64 = 1
	var ival uint64 = 1
	var r uint64 = 2
	
	for i := 0 ; i <= 64 ; i ++ {
		val +=   ival * uint64(math.Pow(float64(r), float64(i)))
	}
	fmt.Printf("val=%v\n", val)
	return math.MaxUint64
	panic("Please implement the Total function")
}
