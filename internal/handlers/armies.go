package handlers

import (
	"encoding/json"

	. "warhammer-aos"
)

type Army struct {
	Army  string `json:"Army"`
	Units []Unit `json:"Units"`
}

type Unit struct {
	Name string `json:"Name"`
	Size string `json:"Size"`
}

func Armies() []Army {
	folder, err := Files.ReadDir("data/json")
	if err != nil {
		panic(err)
	}

	var sections []Army
	for _, file := range folder {
		b, err := Files.ReadFile("data/json/" + file.Name())
		if err != nil {
			panic(err)
		}

		var a Army
		err = json.Unmarshal(b, &a)
		if err != nil {
			panic(err)
		}

		sections = append(sections, a)
	}

	return sections
}
