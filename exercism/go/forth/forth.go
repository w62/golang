// Package forth exports one function: Forth which
// evaluates a very simple Forth expression.
package forth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type stack struct {
	items []int
}

func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// Forth implements a basic Forth expression evaluator.
// Input is a slice of strings. Initially I naively thought that it passes
// everything in a token. Yet after some testings, it probably supplies
// multiple lines of input.
//
// 1. Extract the input into tokens is successful.
func Forth(input []string) ([]int, error) {

	fmt.Printf("raw input=%s rows=%d\n Type is: %-10T", input, len(input), input)

	var (
		x []int
	)

	for _, s := range input {

		//		fmt.Printf("row data=%s row: %d\n Type is: %-10T",s, i, s)
		t := strings.Split(s, " ")

		for _, token := range t {

			//		fmt.Printf("token data=%s len: %d\n Type is: %-10T",token, len(token), token)

			operand, err := strconv.Atoi(token)

			if err == nil { // got an operand
				x = append(x, operand)
				//				fmt.Printf("operand=%d\n", operand)
			} else { // got an operator

				if len(token) == 1 {  // +-*/
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

					case ";":

					default:
						return x, errors.New("Unknown arithemic operator")
					}

				} else { // dup drop swap over
					if len(x) <=0 {
						return x, errors.New("Nothing to Operate. stack underflow")
					}


					//					fmt.Printf("operator=%s\n", token)

					token = strings.ToUpper(token)

					switch token {
					case "DUP":
					if len(x) <1 {
						return x, errors.New("Nothing to Dup. stack underflow")
					} else {
						x = append(x, x[len(x)-1])
					}
					case "DROP":
						if len(x) <1 {
							return x, errors.New("Nothing to drop. stack underflow")
						} else {
							x = x[:len(x)-1]
						}
					case "SWAP":
						if len(x) <2 {
							return x, errors.New("Not enough elements to swap.")
						} else {
							l := len(x) - 1
							m := len(x) - 2
							q := x[l]
							p := x[m]
							x = x [:m]
							x = append(x, q)
							x = append(x, p)
						}
					case "OVER":
						if len(x) <2 {
							return x, errors.New("Not enough elements to swap.")
						} else {
							x = append(x, x[len(x)-2])
						}
					default:
						return x, errors.New("unknown stack operator")

					}

				}

				//				fmt.Printf("operator=%s\n", token)
			}

			//			fmt.Println("token=", token)
		}
	}

	fmt.Printf("x=%v len=%d", x, len(x))
	return x, nil

	panic("Please implement the Forth function")
}
