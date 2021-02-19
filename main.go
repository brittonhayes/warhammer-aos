package main

import (
	"embed"
	"encoding/json"
	"os"
	"text/template"
)

//go:embed data/json/*.json
var f embed.FS

//go:embed templates/json.tmpl
var jsonTemplate string

type Section struct {
	Faction string
	Items   []Mini
}

type Mini struct {
	Name string
	Size string
}

func main() {
	genJSON()
}

func genJSON() {
	folder, err := f.ReadDir("data/json")
	if err != nil {
		panic(err)
	}

	var sections []Section
	for i, file := range folder {
		b, err := f.ReadFile("data/json/" + file.Name())
		if err != nil {
			panic(err)
		}

		var s Section
		err = json.Unmarshal(b, &s)
		if err != nil {
			panic(err)
		}

		if i < 3 {
			sections = append(sections, s)
		}
	}

	tpl, _ := template.New("json").Parse(jsonTemplate)
	if err := tpl.Execute(os.Stdout, sections); err != nil {
		panic(err)
	}
}
