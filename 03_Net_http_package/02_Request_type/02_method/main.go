package main

import (
	"html/template"
	"net/http"
	"log"
	"net/url"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type airplane int

func (a airplane) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Submissions map[string][]string
		Header http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main () {
	var a airplane
	http.ListenAndServe(":8080", a)
}