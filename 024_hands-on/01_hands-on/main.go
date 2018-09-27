package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseGlob("tpl/*.gohtml"))
	data := "This is from dog"
	e := tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if e != nil {
		log.Fatalln(e)
	}
}
