package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// ParseGlob returns (* Template, error) which is what Must takes in
	// basically Must is hanlding the error checking
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Think of *Template as a container for your templates
// Execute can be called to execute container as a whole, it takes in a writer, data interface
// ExecuteTemplate is where we can specify specific templates to execute, takes a writer, string, data interfact
