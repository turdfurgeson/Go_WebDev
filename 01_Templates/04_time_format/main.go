package main

import (
	"html/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01 - 02 - 2006")
}

func main() {

}
