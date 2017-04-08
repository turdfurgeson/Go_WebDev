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

func (p person) GiveMeANumber() int {
	return 4
}

func (p person) DoubleNum() int {
	return p.Age * 2
}

func (p person) PipelineIt(x int) int {
	return 2 * x
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := person{
		Name: "Beastie Boy Mike D",
		Age:  51,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}

}
