package main
import (
	"fmt"
	"regexp"
)


func match (input string) {
	matched, err := regexp.MatchString(`:.*;`, input)

	if err == nil {
		if matched {
		fmt.Printf("%s is command\n", input)
	} else {
		fmt.Printf("%s no command\n", input)
	}
	} else {
		fmt.Printf("syntax error\n")
	}
}

func main() {
	match(": my-own dup over ;")
	match("dup over ;")
	match(":dup over ")

}
