// Exercise 1.4: Modify dup2 to print the names of all files in which each 
// duplicated lines occurs.

// Added another map called filename
// Whenever the line count is larger than 1, add an entry to filename
// finally print out the item at the end.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filename := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)

			f.Close()

			for line, n := range counts {
				if n > 1 {

					if _, found := filename[line]; !found {
						filename[line] = arg
					}
				}
			}

		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t\t%s\n", n, line, filename[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential error from input.Err()
}
