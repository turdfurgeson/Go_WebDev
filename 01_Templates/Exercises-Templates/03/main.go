package main

import (
	"html/template"
	"log"
	"os"
)

type menu struct {
	Course string
	Price  float64
}

type meal struct {
	Time string
	Menu []menu
}

type Meals []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b1 := Meals{
		meal{
			Time: "Breakfast",
			Menu: []menu{
				menu{
					Course: "Scrambled Eggs",
					Price:  4.99,
				},
				menu{
					Course: "Green Eggs and Ham",
					Price:  13.99,
				},
				menu{
					Course: "Chicken Fried Steak with Bacon",
					Price:  17.99,
				},
			},
		},
		meal{
			Time: "Lunch",
			Menu: []menu{
				menu{
					Course: "Club Sandwich",
					Price:  7.99,
				},
				menu{
					Course: "Chicken Soup",
					Price:  8.75,
				},
				menu{
					Course: "Pork Chop with Veggies",
					Price:  12.99,
				},
			},
		},
		meal{
			Time: "Dinner",
			Menu: []menu{
				menu{
					Course: "Steak with frites",
					Price:  19.99,
				},
				menu{
					Course: "Fried Fish with Asaparagus",
					Price:  14.99,
				},
				menu{
					Course: "Shrimp and Chips",
					Price:  23.50,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, b1)
	if err != nil {
		log.Fatalln(err)
	}
}
