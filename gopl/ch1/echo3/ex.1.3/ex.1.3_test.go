// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
//	"time"
	"testing"
)
/*
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
    secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed.\n",secs)
}
*/
func Join_test(b *testing.B) {
	for i:= 0; i < b.N; i++{
	fmt.Println(strings.Join(os.Args[1:], " "))
	}
}

func Dump_test(b *testing.B) {
	for i:= 0; i < b.N; i++{
			s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	}
}
