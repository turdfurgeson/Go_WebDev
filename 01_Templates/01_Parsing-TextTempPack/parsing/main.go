package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parses the folder(glob) passed which returns a pointer to a template
	// and an error. Think of template as a container holding all files which
	// are being parsed.
	tpl, err := template.ParseGlob("template/*")
	if err != nil {
		log.Fatalln(err)
	}
	//Execute take a writer interface(os.Stdout) and we pass in no data.
	// Parses the template above, holds it in"container" and then passes
	// it to stdout (our monitor in this case).
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//continue to add multiple files to pointer to Template("container")
	tpl, err = tpl.ParseFiles("hello.gohtml", "bye.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	//Since multiple templates are available now, use 'ExecuteTemplate"
	// to specify which one!
	err = tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
