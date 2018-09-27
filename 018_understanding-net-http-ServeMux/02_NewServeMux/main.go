package main

import (
	"io"
	"net/http"
)

type hotdog int // Again implements the Handler interface

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

	mux := http.NewServeMux() // returns *ServMmux
	mux.Handle("/dog/", d)    // ( *ServeMux)Handle(string, Handler) <-- this one will handle Path: /dog/*
	mux.Handle("/cat", c)     // <-- will only handle /cat specifically

	http.ListenAndServe(":8080", mux)
}

// BIG NOTE --- *ServeMux ALSO is a Handler (implements its interface), therefor that value can be passed in the place of a Handler
