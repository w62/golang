// Package forth exports one function: Forth which
// evaluates a very simple Forth expression.
package forth

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Forth implements a basic Forth expression evaluator.
// Input is a slice of strings. Initially I naively thought that it passes
// everything in a token. Yet after some testings, it probably supplies
// multiple lines of input.
//
// 1. Extract the input into tokens is successful.
func Forth(input []string) ([]int, error) {

	if len(input) == 0 {
		fmt.Println("nil input")
	} else {
		fmt.Println("some input")
	}

	fmt.Printf("raw input=%s rows=%d\n", input, len(input))

	var (
		x []int
	)

	fmt.Printf("Before calling match my_input=%s rows=%d\n", input, len(input))

	input = match(input)

	fmt.Printf("After calling my_input=%s rows=%d\n", input, len(input))

	for _, s := range input {

		fmt.Printf("row data=%s \n ", s)
		t := strings.Split(s, " ")

		for _, token := range t {

			//		fmt.Printf("token data=%s len: %d\n Type is: %-10T",token, len(token), token)

			operand, err := strconv.Atoi(token)

			if err == nil { // got an operand
				x = append(x, operand)
				//				fmt.Printf("operand=%d\n", operand)
			} else { // got an operator

				if len(token) == 1 { // +-*/
					if len(x) < 2 {
						return x, errors.New("Not enough operands")
					}
					l := len(x) - 1
					m := len(x) - 2

					//					fmt.Printf("operator=%s\n", token)

					q := x[l]
					p := x[m]
					x = x[:m]

					switch token {
					case "+":
						x = append(x, p+q)

					case "-":
						x = append(x, p-q)

					case "*":
						x = append(x, p*q)

					case "/":
						if q == 0 {
							return x, errors.New("divide by zero")
						}
						x = append(x, p/q)

					default:
						return x, errors.New("Unknown arithemic operator")
					}

				} else { // dup drop swap over
					if len(x) <= 0 {
						fmt.Println("x is:", x)
						s := fmt.Sprintf("error %v\n", x)
						return x, errors.New(s)
						//return x, errors.New("Nothing to Operate.  really here? stack underflow")
					}

//					token = strings.ToUpper(token)

					switch token {
					case "DUP":
						if len(x) < 1 {
							return x, errors.New("Nothing to Dup. stack underflow")
						} else {
							x = append(x, x[len(x)-1])
						}
					case "DROP":
						if len(x) < 1 {
							return x, errors.New("Nothing to drop. stack underflow")
						} else {
							x = x[:len(x)-1]
						}
					case "SWAP":
						if len(x) < 2 {
							return x, errors.New("Not enough elements to swap.")
						} else {
							l := len(x) - 1
							m := len(x) - 2
							q := x[l]
							p := x[m]
							x = x[:m]
							x = append(x, q)
							x = append(x, p)
						}
					case "OVER":
						if len(x) < 2 {
							return x, errors.New("Not enough elements to swap.")
						} else {
							x = append(x, x[len(x)-2])
						}
					default:
						return x, errors.New("unknown stack operator")

					}

				}

			}

		}
	}

	fmt.Printf("final x=%v len=%d\n", x, len(x))
	return x, nil

	panic("Please implement the Forth function")
}

func match(input []string) (result []string) {

	var output []string

	lookup := make(map[string]string)

	fmt.Println("input", input)
	for _, s := range input {
		// The regular expression means
		// a line delimited by : and ;
		// the first token must be alphabet with hyphen

		s = strings.ToUpper(s)

		pattern := `^: (?P<marco>[a-zA-Z.+-/*\-]+) (?P<details>.*) ;$`
		cmd := regexp.MustCompile(pattern)

		matched := cmd.FindString(s)

		if matched == "" {

			output = append(output, s)
			continue
		}
		/*

		fmt.Println("matched :", matched)
		fmt.Println("matched len:", len(matched))

		fmt.Println("results shrinked:", results)
*/
		marco := cmd.SubexpIndex("marco")
		details := cmd.SubexpIndex("details")
		results := cmd.FindStringSubmatch(matched)

		leftValue := results[marco]
		rightValue := results[details]

		fmt.Println(lookup)

		if val, ok := lookup[rightValue]; ok {
			delete(lookup, rightValue)
			lookup[rightValue] = val
		fmt.Println("right value found", val)
		} 

		if val, ok := lookup[leftValue]; ok {
			delete(lookup, leftValue)
		fmt.Println("left value deleted", val)
		}
	

		fmt.Println("left", leftValue)
		fmt.Println("right", rightValue)

//		lookup[results[marco]] = results[details]
		lookup[leftValue] = rightValue
	}

	fmt.Println("maps:", lookup)
	fmt.Println("output:", output)

	for token, value := range lookup {
		fmt.Println("token", token)
		switch token {
		case "+":
			token = "\\+"
		case "-":
			token = "\\-"
		case "*":
			token = "\\*"
		case "\\":
		default:
		}
		temp := regexp.MustCompile(token)
		for i, _ := range output {
			output[i] = temp.ReplaceAllString(output[i], value)
		}
	}

	return output

}
