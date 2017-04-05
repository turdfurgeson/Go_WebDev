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
	bands := map[string]string{
		"Psychedelic Rock": "Pink FLoyd",
		"Classic Rock":     "Led Zepplin",
		"80's Rock":        "Guns n Roses",
		"Grunge":           "Nirvana",
	}

	err := tpl.Execute(os.Stdout, bands)
	if err != nil {
		log.Fatalln(err)
	}
}
