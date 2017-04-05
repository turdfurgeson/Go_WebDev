package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	bands := []string{"Pink FLoyd", "The Who", "Led Zepplin", "The Rolling Stones", "Creedence CLearwater Revival"}

	err := tpl.Execute(os.Stdout, bands)
	if err != nil {
		log.Fatalln(err)
	}
}
