package main

import (
	"log"
	"os"
	"strings"
)

func readInput(filename string) [][]string {
	fileContents, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	patterns := [][]string{}

	sections := strings.Split(string(fileContents), "\n\n")
	for _, section := range sections {
		lines := strings.Split(section, "\n")

		// Remove trailing newline
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}

		patterns = append(patterns, lines)
	}

	return patterns
}
