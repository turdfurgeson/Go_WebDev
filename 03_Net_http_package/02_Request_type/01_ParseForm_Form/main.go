package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

//setting type to use as Handler
type airplane int

//implements any airplane type as handler
func (a airplane) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//have to ParseForm BEFORE .form is called
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//.form gives us map[string][]string with key and value pair inside
	//range over these in index.html
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var a airplane
	//takes data and returns it in html with handler a
	http.ListenAndServe(":8080", a)
}
