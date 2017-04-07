package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	food := []string{"pizza", "tacos", "ice cream"}

	err := tpl.Execute(os.Stdout, food)
	if err != nil {
		log.Fatalln(err)
	}
}
