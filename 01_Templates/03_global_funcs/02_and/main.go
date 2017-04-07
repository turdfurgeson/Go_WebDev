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

type band struct {
	Name    string
	Song    string
	Awesome bool
}

func main() {
	pinkFloyd := band{
		Name:    "Pink Floyd",
		Song:    "Wish You Were Here",
		Awesome: true,
	}

	theRollingStones := band{
		Name:    "The Rolling Stones",
		Song:    "Beast of Burden",
		Awesome: true,
	}

	backStreetBoys := band{
		Name:    "Backstreet Boys",
		Song:    "Everybody",
		Awesome: false,
	}

	data := []band{pinkFloyd, theRollingStones, backStreetBoys}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
