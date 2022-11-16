package erratum

import (
//	"errors"
	"fmt"
)

func Use(opener ResourceOpener, input string) error {

	// Open resource
	resource, err := opener()

		fmt.Printf("resource=%T value=%v, err=%T value=%v\n", resource, resource, err, err)

	re, ok := err.(TransientError)

	if ok {
		fmt.Printf("TransientError again")
		fmt.Printf(" re=%T\n", re)
	
	} 
	// close the resource
		defer resource.Close()
	// Call Frob(input) upon the opened resource
	resource.Frob(input)

	return nil
	panic("Please implement the Use function")
}
