package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("starting-files/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.HandleFunc("/apply/", apply)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	e := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}
}

func about(w http.ResponseWriter, req *http.Request) {
	e := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}
}

func contact(w http.ResponseWriter, req *http.Request) {
	e := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		e := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		if e != nil {
			log.Fatalln(e)
		}
		return
	}
	e := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}
}
