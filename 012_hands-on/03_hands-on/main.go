package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "The Trevor Inn",
			Address: "123 Go Way",
			City:    "Dallas",
			Zip:     "75014",
			Region:  "DFW",
		},
		hotel{
			Name:    "MicroService Hotel",
			Address: "333 Big Street",
			City:    "Theodore",
			Zip:     "36582",
			Region:  "AL",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
