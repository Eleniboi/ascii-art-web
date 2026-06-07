package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ArtBuilder(text, banner string) (string, error) {

	text = strings.ReplaceAll(text, "\r\n", "\n")
	file, err := os.Open("banners/"+banner)

	if err != nil {

		return "", fmt.Errorf("error reading file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines strings.Builder

	for scanner.Scan() {
		//this line throw away "\n" so i have to manually add a newline while building
		lines.WriteString(scanner.Text())
		lines.WriteString("\n")

	}
	err = scanner.Err()

	if err != nil {
		return "", fmt.Errorf("error scanning file: %w", err)
	}
	//the builder got stored in another variable
	line := lines.String()

	fileline := strings.Split(line, "\n")
	fileline = fileline[1:]

	var build strings.Builder

	splitinput := strings.Split(text, "\n")
	for _, word := range splitinput {
		for row := 0; row < 8; row++ {
			for _, ch := range word {

				start := int(ch-32) * 9
				end := start + 8
				build.WriteString(fileline[start:end][row])
			}
			build.WriteString("\n")
		}
	}
	return build.String(), nil
}