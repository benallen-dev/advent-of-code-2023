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
		patterns = append(patterns, lines)
	}

	return patterns
}
