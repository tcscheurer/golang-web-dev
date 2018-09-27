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

	// type http.HandlerFunc's underlying type is func(http.ResponseWriter, *http.Request)
	// looks like this in docs --> type http.HandlerFunc func(http.ResponseWriter, *http.Request)

	http.Handle("/dog", http.HandlerFunc(d)) // here we are passing func(http.ResponseWriter, *http.Request) into HandlerFunc
	http.Handle("/cat", http.HandlerFunc(c)) // which returns a value of type Handler, which implements the Handler interface
	// which is the implementation of the method ServeHttp(http.ResponseWriter, *http.Request)

	http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L
