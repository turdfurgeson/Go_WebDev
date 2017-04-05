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

type player struct {
	Name string
	Era  string
}

type items struct {
	Name []band
	Type []player
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

	cassette := player{
		Name: "cassette",
		Era:  "1980s",
	}

	cd := player{
		Name: "compact disc",
		Era:  "1990s",
	}

	vinyl := player{
		Name: "vinyl",
		Era:  "1970s",
	}

	bands := []band{pinkFloyd, ledZepplin, rollingStones}
	players := []player{cassette, cd, vinyl}

	data := items{
		Name: bands,
		Type: players,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
