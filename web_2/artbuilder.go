package main

import (
	"log"
	"os"
	"strings"
)

func Artbuilder(text, banner string) (string, error) {

	text = strings.ReplaceAll(text, "\r\n", "\n")

	file, err := os.ReadFile("banner/" + banner)

	if err != nil {
		log.Fatal(err)
	}

	fileLine := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	fileLine = fileLine[1:]

	var build strings.Builder
	splitText := strings.Split(text, "\n")

	for _, word := range splitText {

		for row := 0; row < 8; row++ {
			for _, ch := range word {
				start := int(ch-32) * 9
				end := start + 8

				build.WriteString(fileLine[start:end][row])
			}
			build.WriteString("\n")
		}
	}
	return build.String(), nil
}
