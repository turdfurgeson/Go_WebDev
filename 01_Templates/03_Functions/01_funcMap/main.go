package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

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

	//composite literal format. Define struct and pass structs on fly.
	data := struct {
		Name []band
		Type []player
	}{
		bands,
		players,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
