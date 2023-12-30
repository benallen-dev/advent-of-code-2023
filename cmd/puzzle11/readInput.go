package main

import (
	"log"
	"os"
	"strings"
	"regexp"
)

var galaxyRegex = regexp.MustCompile(`#`)

func transpose(input []string) []string {
	output := strings.Split(input[0], "")

	for i, line := range input {
		if i == 0 {
			continue
		}

		for j, char := range line {
			output[j] += string(char)
		}
	}

	return output
}

func duplicateEmptyLines(input []string) []string {
	output := []string{}

	// empty lines are easy
	for _, line := range input {
		if galaxyRegex.MatchString(line) {
			output = append(output, line)
			continue
		}

		// We have a line with no galaxies, so duplicate it
		output = append(output, line)
		output = append(output, line)
	}

	return output
}

func expandUniverse (input []string) []string {
	// First, duplicate empty lines
	output := duplicateEmptyLines(input)
	
	// Cool, but now we need to deal with columns
	output = transpose(output)
	output = duplicateEmptyLines(output)

	// Now transpose it back to the right shape
	output = transpose(output)

	return output
}

func readInput(filename string) []string {
	fileContents, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileContents), "\n")
	lines = lines[:len(lines)-1] // Remove the last line because it's empty

	// Expand the universe
	lines = expandUniverse(lines)

	return lines
}
