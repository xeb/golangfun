package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name string
}

func (p Person) Label() string {
	return "This is " + string(p.Name)
}

func main() {
	tmpl, err := template.New("").Parse(`{{.Label}}`)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	tmpl.Execute(os.Stdout, &Person{Name: "Bob"})
}
