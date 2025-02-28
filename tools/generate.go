package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	perpetua "github.com/perpetuatheme/go"
)

//go:generate go run generate.go

func main() {
	name := "alacritty.tmpl.toml"
	t := template.Must(template.New(name).ParseFiles(name))
	generate(t, perpetua.Light)
	generate(t, perpetua.Dark)
}

func generate(t *template.Template, p perpetua.Palette) {
	f, err := os.Create("../perpetua-" + strings.ToLower(p.Name()) + ".toml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = t.Execute(f, p)
	if err != nil {
		log.Fatal(err)
	}
}
