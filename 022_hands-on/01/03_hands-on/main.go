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
	err := tpl.ExecuteTemplate(res, "dog.gohtml", "Hot Dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func name(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "name.gohtml", "Trevor Page")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", name)

	http.ListenAndServe(":8080", nil)
}
