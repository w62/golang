package thefarm

import (
	"fmt"
	"errors"
)

// See types.go for the types defined for this exercise.
type SillyNephewError struct {
	cows int
//	msg string
Err error
}
func (e *SillyNephewError) Error() string {
		return_msg := fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows) 
	return return_msg
}

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	var ne  SillyNephewError
	i, err := weightFodder.FodderAmount()
	if i <0 && err == ErrScaleMalfunction {
		return 0, errors.New("negative fodder")
	}

	if i < 0 && err == nil {
		return 0, errors.New("negative fodder")
	}
	if i == 0 {
		if err == nil {
		return 0, err 
		}
		return 0, errors.New("division by zero") 
	}

	if cows == 0{
		return 0, errors.New("division by zero") 
	}

	if cows < 0 {
		ne.cows = -10
		return 0, &SillyNephewError {
			cows: cows,
			Err: errors.New(""),
		}
	}
	if err == nil {
		if err == ErrScaleMalfunction {
			fmt.Println("ErrScaleMalfunction")
			i *= 2.0
		}
	return i / float64(cows), nil
} else if err == ErrScaleMalfunction {
			i *= 2.0
	return i / float64(cows), nil
		} else if err != nil {
	return 0, err
}
	panic("Please implement DivideFood")
}
