package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // <- method from http.Request type
	// REALLY IMPORTANT ^ This method MUST be called in order to Access req.Form (or req.PostForm) like on line 18
	// Also, req.Form contains all query params as well as ones in the body of request
	// Req.PostForm will only contain the ones in the body
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
