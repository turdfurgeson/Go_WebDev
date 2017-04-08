package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

type guitarist struct {
	person
	Style string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	g1 := guitarist{
		person: person{
			Name: "David Gilmour",
			Age:  71,
		},
		Style: "blues",
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}

}
