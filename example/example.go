package main

//go:generate templify ex.tpl

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	//Version without embedded template. Reading from file.
	//tpl, err := template.ParseFiles("ex.tpl")

	//Embedded template
	tpl, err := template.New("ex").Parse(exTemplate())

	if err != nil {
		fmt.Printf("Error parsing code generation template\n%v\n", err)
		os.Exit(1)
	}

	data := struct {
		Head    string
		Title   string
		Content string
	}{
		Head:    "<meta stuff>",
		Title:   "Supertitle",
		Content: "This could be your content",
	}

	tpl.Execute(os.Stdout, data)

}
