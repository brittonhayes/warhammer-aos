package warhammer_aos

import (
	"embed"
	"encoding/json"
)

//go:embed data/json/*.json
var f embed.FS

//go:embed templates/json.tmpl
var JsonTemplate string

type Section struct {
	Army  string `json:"Army"`
	Units []Unit `json:"Units"`
}

type Unit struct {
	Name string `json:"Name"`
	Size string `json:"Size"`
}

func GenJSON() []Section {
	folder, err := f.ReadDir("data/json")
	if err != nil {
		panic(err)
	}

	var sections []Section
	for _, file := range folder {
		b, err := f.ReadFile("data/json/" + file.Name())
		if err != nil {
			panic(err)
		}

		var s Section
		err = json.Unmarshal(b, &s)
		if err != nil {
			panic(err)
		}

		sections = append(sections, s)
	}

	return sections
}
