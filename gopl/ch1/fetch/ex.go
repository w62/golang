// Fetch prints the content found at a URL.
package main

import (
	"fmt"
//	"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main () {
	for _, url := range os.Args[1:] {

		// Exercise 1.8: Modify fetch to add the prefix http:// to each argument
		// URL if it is missing. You might want to use strings.HasPrefix.

		if !strings.HasPrefix(url, "http://") {
		fmt.Println("http:// not found")
		url = "http://" + url
	}
		resp, err := http.Get(url)


		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Exercise 1.9: Modify fetch to also print the HTTP status code,
		// found int resp.Status.
		fmt.Println("HTTP status code:", resp.Status)

//		b, err := ioutil.ReadAll(resp.Body)
// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes 
// to dst. Use it instead of ioutil.ReadAll to copy the response body to os.
// Stdout without requiring buffer large enough to hold the entire stream.
// Be sure to check the error result of io.Copy.

		b, err := io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "io.Copy: %v\n", err)
			os.Exit(1)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}



