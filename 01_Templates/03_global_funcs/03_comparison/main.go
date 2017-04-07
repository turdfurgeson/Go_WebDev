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
	Songs := struct {
		Hits   int
		Albums int
	}{
		9,
		17,
	}

	err := tpl.Execute(os.Stdout, Songs)
	if err != nil {
		log.Fatalln(err)

	}

}
