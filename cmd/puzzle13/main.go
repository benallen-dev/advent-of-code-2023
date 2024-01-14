package main

import (
	"log"
//	"fmt"

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
func checkReflection(pattern []string, startLine int) (linePairs [][]int, isReflection bool) {
	distanceToEnd := len(pattern) - startLine
	iterations := min(startLine+1, distanceToEnd-1)

		log.Println("Length:", len(pattern))
		log.Println("Start line:", startLine)
		log.Println("Distance to end:", distanceToEnd)
		log.Println("Iterations:", iterations)

	// Check each line pair for equality
	for i := 0; i < iterations; i++ {
		log.Println("Checking", startLine-i, "and", startLine+1+i)

		if pattern[startLine-i] != pattern[startLine+1+i] {
			log.Println("Not equal:", pattern[startLine-i], pattern[startLine+1+i])
			return [][]int{}, false // If any pair is not equal, the pattern is not a reflection
		}

		linePairs = append(linePairs, []int{startLine - i, startLine + 1 + i})
	}

	return linePairs, true
}

func findReflection(pattern []string, printResult bool, isTransposed bool) (line int, ok bool) {
	line, ok = findDoubleLine(pattern)

	if ok {
		log.Println("Found double line at", line)
		if linePairs, isReflection := checkReflection(pattern, line); isReflection {

			if printResult {
				if isTransposed {
					printPatternTransposed(pattern, linePairs)
				} else {
					printPattern(pattern, linePairs)
				}
			}

			// Because the lines are 1-indexed we need to add 1 to our array index
			return line + 1, true
		}
	}

	return 0, false
}

func main() {
	log.SetPrefix(color.Green + "[ # 13 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("edgecases.txt")

	totalHorizontal := 0
	totalVertical := 0

	for i, line := range input {
		log.Println()
		log.Println("Pattern:", i)
		horizontalLine, okHorizontal := findReflection(line, true, false)
		verticalLine, okVertical := findReflection(strutil.Transpose(line), true, true)

		// Find a reflection
		// Print the pattern as given with the lines/columns highlighted

		// One by one check if it makes sense

		if okHorizontal && okVertical {
			log.Println("Both horizontal and vertical reflections found, qouiiiii")
		}

		if okHorizontal {
			totalHorizontal += horizontalLine

			log.Printf("Adding %d to horizontal count", horizontalLine)
			// log.Println("Horizontal reflection found on line", horizontalLine)
		}

		if okVertical {
			totalVertical += verticalLine

			log.Printf("Adding %d to vertical count", verticalLine)
			// log.Println("Vertical reflection found on line", verticalLine)
		}

		if !okHorizontal && !okVertical {
			// quoiiii
			printPattern(line, [][]int{})
//			printPattern(strutil.Transpose(line), [][]int{})
		}

// 		log.Println("Press enter to continue")
// 		fmt.Scanln()
	}

	total := (100 * totalHorizontal) + totalVertical

	log.Println("Total:", total)
}
