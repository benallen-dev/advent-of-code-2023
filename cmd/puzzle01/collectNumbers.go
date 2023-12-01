package main

import (
	"regexp"
)

// Global regexes so we're sure they only get compiled once
var reNotNumber = regexp.MustCompile(`[^0-9]`)

func collectNumbers(line string) string {
	return reNotNumber.ReplaceAllLiteralString(line, "")
}

