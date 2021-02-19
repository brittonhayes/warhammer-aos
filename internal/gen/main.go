package main

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

//go:embed data/*.csv
var f embed.FS

func main() {
	folder, err := f.ReadDir("data/")
	if err != nil {
		panic(err)
	}

	for _, file := range folder {
		b, err := f.ReadFile("data/" + file.Name())
		if err != nil {
			panic(err)
		}

		r := csv.NewReader(strings.NewReader(string(b)))
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(record)
		}
	}

}
