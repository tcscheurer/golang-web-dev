package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog", d) // Handle func does not take in a Handler as second param
	http.HandleFunc("/cat", c) // it takes a func(http.ResponseWriter, *http.Request) instead

	http.ListenAndServe(":8080", nil)
}
