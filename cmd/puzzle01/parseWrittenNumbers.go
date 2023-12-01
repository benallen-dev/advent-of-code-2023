package main

import (
	"log"
	"regexp"
)

var regexWrittenNumber = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|zero)`)

// Global map so again, not evaluating this every time we call parseLine, basically
var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func replaceWithDigit(writtenNumber string) string {

	newvalue, exists := numberMap[writtenNumber]
	if !exists {
		log.Panic("Tried accessing non-existent map property:", writtenNumber)
	}

	return newvalue
}

func parseWrittenNumbers(line string) string {
	
	// for i := 0; i < len(line); i++ {

	// }

	return regexWrittenNumber.ReplaceAllStringFunc(line, replaceWithDigit)
}

