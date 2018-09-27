package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog", d) // <- Since we are calling the package level function Handle instead of the one that takes *ServeMux as a receiver, we are registering these with the default servemux
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil) // <- nil tells the compiler to use the Default Serve Mux
}
