package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func home(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "home.gohtml", "Home Page")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "home.gohtml", "Hot Dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func name(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "home.gohtml", "Trevor Page")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(name))

	http.ListenAndServe(":8080", nil)
}
