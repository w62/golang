package main

import (
	"fmt"
	"io"
//	"io/ioutil"
	"net/http"
	"os"
	"time"
	"math/rand"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	// Exercise 1.10: Find a web site that produces a large amount of data.
	// Investigate caching by running fetchall twice in succession to see 
	// whether the reported time changes much. Do you get the same content each
	// time? Modify fetchall to print its output to a file so that it can be 
	// examined.

	// Exercise 1.11: Try fetchall with longer arguments list, such as samples 
	// from the top million web sites available at alexa.com. Dead now!! How
	// does the program behave if a web site just don't respond? (Section 8.9
	// describes mechanisms for coping in such cases.)

	// Use date and a random number for the file names.

	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	filename := fmt.Sprintf("%v-%v.txt", start.Format("2006-01-02-15:04:05"),
	rand.Intn(99 - 0))
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %v\n", err)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// inbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Copy error: %v\n", err)
		return
	}
 if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "File close error: %v\n", err)
		return
        }
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
