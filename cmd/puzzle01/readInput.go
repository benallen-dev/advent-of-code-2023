package main

import (
	"log"
	"os"
	"strings"
)

func readInput(filename string) []string {
	fileBytes, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	return strings.Split(string(fileBytes), "\n")
}
