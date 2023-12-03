package main

import (
	"log"
	"regexp"
	"strconv"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/math"
)

var regexNumber = regexp.MustCompile(`[0-9]+`)
var regexSymbol = regexp.MustCompile(`[^.0-9]`)

func partOne(lines []string) int {
	total := 0
	for lineNumber, line := range lines {
		// Part 1
		numbers := regexNumber.FindAllStringIndex(line, -1)

		for _, numberIdx := range numbers {
			// Get numbers adjecent to symbols
			indexLeft := math.Clamp(numberIdx[0]-1, 0, len(line)-1)
			indexRight := math.Clamp(numberIdx[1]+1, 0, len(line)-1)

			prevLineNumber := math.Clamp(lineNumber-1, 0, len(lines)-1)
			nextLineNumber := math.Clamp(lineNumber+1, 0, len(lines)-1)

			// For lines 0 and len(lines)-1, there is no line before/after so we just use the current line
			lineBefore := lines[prevLineNumber][indexLeft:indexRight]
			lineCurrent := lines[lineNumber][indexLeft:indexRight]
			lineAfter := lines[nextLineNumber][indexLeft:indexRight]

			if regexSymbol.MatchString(lineBefore) || regexSymbol.MatchString(lineCurrent) || regexSymbol.MatchString(lineAfter) {
				// Extract the number from the line
				number, err := strconv.Atoi(line[numberIdx[0]:numberIdx[1]])
				if err != nil {
					log.Fatal(err)
				}

				total += number
			}
		}
	}

	return total
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 03 ] " + color.Reset)

	lines := readInput("input.txt")

	totalParts := partOne(lines)

	// Find all numbers
	// if completely surrounded by '.', continue
	// otherwise add number to running total

	log.Println("Total part 1:", totalParts)
}
