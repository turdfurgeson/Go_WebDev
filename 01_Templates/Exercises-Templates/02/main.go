package main

import (
	"html/template"
	"log"
	"os"
)

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h1 := Hotel{
		Name:    "Holiday Inn",
		Address: "1234 Skydrive Ln",
		City:    "Sacramento",
		Zip:     12345,
		Region:  "Northern",
	}

	h2 := Hotel{
		Name:    "Best Western",
		Address: "4567 Tantigold St",
		City:    "San Diego",
		Zip:     57480,
		Region:  "Southern",
	}

	h3 := Hotel{
		Name:    "The Roosevelt",
		Address: "2785 Hollywood Blvd",
		City:    "Los Angeles",
		Zip:     90210,
		Region:  "Central",
	}

	Hotels := []Hotel{h1, h2, h3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
