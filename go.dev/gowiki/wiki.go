package main
// Writing Web Applications
//Up to here:
// https://go.dev/doc/articles/wiki/#tmp_2
// Can write a file to disk at OS level
// Now here https://go.dev/doc/articles/wiki/#tmp_4
// The bare web server

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

