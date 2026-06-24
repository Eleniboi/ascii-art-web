package main

import (
	"log"
	"os"
	"strings"
)

func Artbuilder(text, banner string) (string, error) {


	text = strings.ReplaceAll(text,  "\r\n", "\n")
	file, err := os.ReadFile(banner)

	if err != nil {
		log.Fatal(err)
	}

	fileLine := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	fileLine = fileLine[1:]

	splitInput := strings.Split(text, "\n")

	for _, word := range splitInput{

		for row := 0; row < 8; row++{
			for _, ch := range word{

				startindx := int(ch - 32) *9
				endindx := startindx+8

				
			}
		}
	}
}
