package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world %s!", r.URL.Path[1:])
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sb %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
