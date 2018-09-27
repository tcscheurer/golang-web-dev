package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}

// Any type that defines ServeHTTP(w http.ResponseWriter, req *http.Request) implements the Handler Interface!!!!
// ResponseWriter Implements the writer Interface -> Write([]byte)(int, error)
