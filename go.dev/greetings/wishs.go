package greetings

import (
	"fmt"
)

	func GoodLuck(name string) string {
		p := people {
			name: "Dora", 
			age: 18,
		}
		fmt.Println(p.name, " ", p.age)
		return DefaultMessage + name
	}
