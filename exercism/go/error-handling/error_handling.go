package erratum
// https://go.dev/play/p/N73GzIfloxD 

import (
	//		"errors"
	"fmt"
)

func Use(opener ResourceOpener, input string) error {

	// Open resource
	resource, err := opener()

	// Assuming we have not encountered TransientError to start with.
	terr := false

	for !terr {

		switch err.(type) {
		case nil:
			fmt.Printf("nil err case \n")
			resource.Frob(input)
			terr = true
		case TransientError:
			fmt.Printf("TransientError case %T\n", err)
			resource, err = opener()
		default:
			return fmt.Errorf("too awesome")
		}
	}
	// defer close the resource
	defer func() {
		fmt.Printf("oh shit! panic!")
		if r := recover(); r != nil {
			fmt.Printf("recovered %d type=%T\n", r, r)

		} else {
			fmt.Printf("ur ....recovered %d type=%T\n", err, err)
			fmt.Println("ur....")
		}
		defer resource.Close()
		//		resource.Defrob("moo")
	}()

	fmt.Printf("resource type=%T value=%v, err=%T value=%v\n", resource, resource, err, err)

	switch resource.(type) {
	case nil:
		fmt.Printf("nil resource case \n")
	case mockResource:
		fmt.Printf("mockResource case \n")
	default:
		return fmt.Errorf("other error")

	}

	// Call Frob(input) upon the opened resource

	return nil
	panic("Please implement the Use function")
}
