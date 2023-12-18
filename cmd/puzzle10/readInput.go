package main

import (
	"log"
	"os"
	"strings"
)

func readInput(filename string) []string {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileContents), "\n")
	lines = lines[:len(lines)-1] // Remove the last line because it's empty

	return lines
}
