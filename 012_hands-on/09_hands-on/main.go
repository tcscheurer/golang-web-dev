package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

type csvDataType struct {
	Date   string
	Open   string
	High   string
	Low    string
	Close  string
	Volume string
	Adj    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	csvFile, _ := os.Open("table.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var csvData []csvDataType
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		csvData = append(csvData, csvDataType{
			Date:   line[0],
			Open:   line[1],
			High:   line[2],
			Low:    line[3],
			Close:  line[4],
			Volume: line[5],
			Adj:    line[6],
		})
	}
	err := tpl.Execute(res, csvData)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
