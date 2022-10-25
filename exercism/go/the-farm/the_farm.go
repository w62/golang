package thefarm

import (
	"fmt"
	"errors"
)

// See types.go for the types defined for this exercise.
type SillyNephewError struct {
	cows int
	msg string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}


// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	i, err := weightFodder.FodderAmount()
	if i == 0 {
		return 0, errors.New("division by zero") 
	}

	if cows <= 0 {
		return 0, SillyNephewError.Error()
	}
	fmt.Printf("ErrScale i=%.2f cows=%d cows as float = %.2f err=%v\n", i, 
	cows, float64(cows), err)
	if err == nil {
		if err == ErrScaleMalfunction {
			fmt.Println("ErrScaleMalfunction")
			i *= 2.0
		}
	fmt.Printf("i=%.2f cows=%d cows as float = %.2f\n", i, cows, float64(cows))
	return i / float64(cows), nil
} else if err == ErrScaleMalfunction {
			fmt.Println("ErrScaleMalfunction")
			i *= 2.0
	return i / float64(cows), nil
		} else if err != nil {
	fmt.Println("error")
	return 0, err
}
	panic("Please implement DivideFood")
}
