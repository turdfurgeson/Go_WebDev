package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type band struct {
	Name string
	Song string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	pinkFloyd := band{
		Name: "Pink Floyd",
		Song: "Comfortably Numb",
	}

	err := tpl.Execute(os.Stdout, pinkFloyd)
	if err != nil {
		log.Fatalln(err)
	}
}
