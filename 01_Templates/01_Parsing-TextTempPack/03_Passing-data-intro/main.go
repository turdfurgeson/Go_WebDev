package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main () {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Be the change " +
		"you want to see in the world`)
	if err != nil {
		log.Fatalln(err)
	}
}
