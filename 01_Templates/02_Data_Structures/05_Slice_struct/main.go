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
		Song: "Wish You Were Here",
	}

	ledZepplin := band{
		Name: "Led Zepplin",
		Song: "Stairway to Heaven",
	}

	rollingStones := band{
		Name: "The Rolling Stones",
		Song: "Beast of Burden",
	}

	bands := []band{pinkFloyd, ledZepplin, rollingStones}

	err := tpl.Execute(os.Stdout, bands)
	if err != nil {
		log.Fatalln(err)
	}
}
