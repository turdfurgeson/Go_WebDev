package main

import (
	"html/template"
	"log"
	"os"
)

//defines tpl as pointer to template Template at package scope
var tpl *template.Template

func init() {
	//.Must does error checking for us and we are passing in Parse glod
	// which returns an error as well. What parsGlob returns is exactly
	// what Must takes!
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "hiFelicia.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "byeFelicia.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
