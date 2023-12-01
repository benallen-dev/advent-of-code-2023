package main

import (
	"regexp"
)

var regexWrittenNumber = regexp.MustCompile(`^(one|two|three|four|five|six|seven|eight|nine|zero)`)
var regexNumber = regexp.MustCompile(`[0-9]`)

// Global map so again, not evaluating this every time we call parseLine, basically
var numberMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func parseWrittenNumbers(line string) string {
	output := ""

	// Sliding window from i to end of input. If number, add to
	// output. If written number, convert to number and add to output.
	for i := 0; i < len(line); i++ {
		window := line[i:]

		if regexNumber.MatchString(string(window[0])) {
			output += string(window[0])
			continue
		}

		if regexWrittenNumber.MatchString(window) {
			output += numberMap[regexWrittenNumber.FindString(window)]
		}
	}

	return output
}
