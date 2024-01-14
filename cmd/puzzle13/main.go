package main

import (
	"fmt"
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

func findDoubleLine(pattern []string) (line int, ok bool) {
	for i := 0; i < len(pattern)-1; i++ { // for each line
		if pattern[i] == pattern[i+1] {
			return i, true
		}
	}

	return 0, false
}

// Columns are the same, only transposed, and strings are easier to compare
func checkReflection(pattern []string, startLine int) bool {
	distanceToEnd := len(pattern) - startLine
	iterations := min(startLine+1, distanceToEnd -1)

	// Check each line pair for equality
	for i := 0; i < iterations; i++ {
		if pattern[startLine-i] != pattern[startLine+1+i] {
			return false // If any pair is not equal, the pattern is not a reflection
		}
	}

	return true
}

// There's some funky off-by-one stuff going on here but I'm too tired to try to figure out what it is
func getLinesThatShouldBeEqual(pattern []string, startLine int) (lines [][]int) {
	distanceToEnd := len(pattern) - startLine
	iterations := min(startLine, distanceToEnd)

	for i := 0; i < iterations; i++ {
		lines = append(lines, []int{startLine - i, startLine + 1 + i})
	}

	return lines
}

func findReflection(pattern []string, printResult bool) (line int, ok bool) {
	line, ok = findDoubleLine(pattern)
	if ok && checkReflection(pattern, line) {

		if printResult {
			linePairs := getLinesThatShouldBeEqual(pattern, line)
			printPattern(pattern, linePairs)
		}

		// Because the lines are 1-indexed we need to add 1 to our array index
		return line + 1, true
	}

	return 0, false
}

func printPattern(pattern []string, matchingLines [][]int) {
	colors := []string{color.Red, color.Green, color.Yellow, color.Blue, color.Purple, color.Cyan}
	colorMap := map[int]string{}

	// Get colour map
	for i, linePair := range matchingLines {
		colorMap[linePair[0]] = colors[i]
		colorMap[linePair[1]] = colors[i]
	}

	log.Println("  123456789")
	for i, line := range pattern {
		if lineColor, ok := colorMap[i]; ok {
			log.Println(fmt.Sprintf("%d ", i) + lineColor + line + color.Reset)
		} else {
			log.Println(fmt.Sprintf("%d ", i) + line)
		}
	}
}

func main() {
	log.SetPrefix(color.Green + "[ # 13 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")

	total := 0

	for i, line := range input {
		horizontalLine, okHorizontal := findReflection(line, false)
		verticalLine, okVertical := findReflection(strutil.Transpose(line), false)

		if okHorizontal {
			total += 100 * horizontalLine
		}

		if okVertical {
			total += verticalLine
		}

		if (okHorizontal || okVertical) {
			log.Println()
			log.Println("Pattern:", i)

			if okHorizontal {
				log.Println("Horizontal reflection found on line", horizontalLine)
			}

			if okVertical{
				log.Println("Vertical reflection found on line", verticalLine)
			}

			log.Println()
			log.Println("--------------------")
		}
	}

	log.Println("Total:", total)
}
