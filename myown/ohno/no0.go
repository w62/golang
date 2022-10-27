package unknown

import (
	"fmt"
)

var Umessage0 string

func Um () {
	fmt.Println(Umessage0)
	fmt.Println(Umessage1)
}
func Update1 (newMessage string) string{
	Umessage1 = newMessage
	return "Umessage1 updated to "+Umessage1
}

