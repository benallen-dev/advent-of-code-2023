package main

import (
	"log"
	"os"
	"strings"
)

func readInput(filename string) []string {
	fileContents, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	steps := strings.Split(string(fileContents), ",")

	for i, step := range steps {
		steps[i] = strings.TrimSuffix(step, "\n")
	}


	return steps
}
