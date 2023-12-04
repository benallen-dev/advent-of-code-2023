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
var regexStart = regexp.MustCompile(`\*`)

var DEBUG = false

func partOne(lines []string) int {
	total := 0
	for lineNumber, line := range lines {
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

func partTwo(lines []string) int {

	totalGears := 0

	for lineNumber, line := range lines {

		var numbersAbove [][]int
		var numbersBelow [][]int

		numbersCurrent := regexNumber.FindAllStringIndex(line, -1)

		if lineNumber > 0 {
			numbersAbove = regexNumber.FindAllStringIndex(lines[lineNumber-1], -1)
		}

		if lineNumber < len(lines)-1 {
			numbersBelow = regexNumber.FindAllStringIndex(lines[lineNumber+1], -1)
		}

		stars := regexStart.FindAllStringIndex(line, -1)
		for _, star := range stars {
			var adjacentNumbers []int = []int{}

			indexLeft := math.Clamp(star[0]-1, 0, len(line)-1)
			indexRight := math.Clamp(star[1], 0, len(line)-1)

			for _, numberCurrent := range numbersCurrent {

				// immediately before
				if numberCurrent[1] == star[0] {
					stringNumber := line[numberCurrent[0]:numberCurrent[1]]
					number, err := strconv.Atoi(stringNumber)
					if err != nil {
						log.Fatal(err)
					}

					adjacentNumbers = append(adjacentNumbers, number)
				}

				// immediately after
				if numberCurrent[0] == star[1] {
					stringNumber := line[numberCurrent[0]:numberCurrent[1]]
					number, err := strconv.Atoi(stringNumber)
					if err != nil {
						log.Fatal(err)
					}

					adjacentNumbers = append(adjacentNumbers, number)
				}
			}

			// for each pair in numbersAbove check if any range overlaps with [indexLeft:indexRight]
			for _, numberAbove := range numbersAbove {

				if math.RangeOverlap(numberAbove[0], numberAbove[1]-1, indexLeft, indexRight) {
					stringNumber := lines[lineNumber-1][numberAbove[0]:numberAbove[1]]
					number, err := strconv.Atoi(stringNumber)
					if err != nil {
						log.Fatal(err)
					}

					adjacentNumbers = append(adjacentNumbers, number)
				}
			}

			// Now do the same for numbersBelow
			for _, numberBelow := range numbersBelow {

				if math.RangeOverlap(numberBelow[0], numberBelow[1]-1, indexLeft, indexRight) {
					stringNumber := lines[lineNumber+1][numberBelow[0]:numberBelow[1]]
					number, err := strconv.Atoi(stringNumber)
					if err != nil {
						log.Fatal(err)
					}

					adjacentNumbers = append(adjacentNumbers, number)
				}
			}

			if DEBUG {
				if lineNumber > 0 && lineNumber < len(lines)-1 {
					log.Println(lines[lineNumber-1][star[0]-3 : star[1]+3])
					log.Println(line[star[0]-3 : star[1]+3])
					log.Println(lines[lineNumber+1][star[0]-3 : star[1]+3])
				}
				log.Println("Adjecent numbers:", adjacentNumbers)
			}

			if len(adjacentNumbers) == 2 {
				totalGears += (adjacentNumbers[0] * adjacentNumbers[1])
			} else if DEBUG {
				log.Printf("Discarding %v", adjacentNumbers)
			}

		}
	}

	return totalGears
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 03 ] " + color.Reset)

	lines := readInput("input.txt")

	// I'm aware I'm iterating over the lines twice, but O(2n) is still O(n) and there's only 140 lines
	totalParts := partOne(lines)
	totalGears := partTwo(lines)

	log.Println("Total part 1:", totalParts)
	log.Println("Total part 2:", totalGears)
}
