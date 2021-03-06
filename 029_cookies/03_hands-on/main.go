package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/clear", clear)
	http.HandleFunc("/", handleIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("my-counter")
	if e != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "my-counter",
			Value: "1",
		})
		fmt.Fprintln(w, "Cookie-Counter val is: ", 1)
		return
	}

	val, e := strconv.Atoi(c.Value)
	if e != nil {
		log.Println(e)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "my-counter",
		Value: strconv.Itoa(val + 1),
	})
	count, e := req.Cookie("my-counter")
	if e != nil {
		log.Println(e)
	}
	fmt.Fprintln(w, "Cookie-Counter val is: ", count.Value)
}

func clear(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("my-counter")
	if e != nil {
		log.Println("Clear request came in but cookie doesn't exist")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
