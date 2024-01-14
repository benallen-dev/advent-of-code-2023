package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

func findDoubleLines(pattern []string) (lines [][]int) {
	for i := 0; i < len(pattern)-1; i++ { // for each line

		differences := countDifferences(pattern, []int{i, i + 1})
		// We want to count lines that have 1 difference for part 2
		if differences < 2 {
			lines = append(lines, []int{i, i + 1})
		}
	}

	return lines
}

func generateLinePairs(pattern []string, doubleLines []int) (linePairs [][]int) {
	left := doubleLines[0]
	right := doubleLines[1]

	distanceToEnd := len(pattern) -1 - right
	iterations := min(left, distanceToEnd)

	for i := 0; i <= iterations; i++ {
		linePairs = append(linePairs, []int{left - i, right + i})
	}

	return linePairs
}

func countDifferences(pattern []string, linePair []int) (differences int) {
	differences = 0 // Not sure if init is needed but let's be explicit

	a := pattern[linePair[0]]
	b := pattern[linePair[1]]

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}

	return differences
}


func processPattern(pattern []string) (noDifferenceLines int, singleDifferenceLines int) {

	oldPrefix := log.Prefix()
	log.SetPrefix(color.Cyan + "[ # 13 :: processPattern ] " + color.Reset)

	// find all sets of double lines
	doubleLines := findDoubleLines(pattern)

	reflectionCandidates := [][][]int{}

	// for each set of double lines, generate a list of line pairs
	for _, doubleLine := range doubleLines {
		linePairs := generateLinePairs(pattern, doubleLine)
		reflectionCandidates = append(reflectionCandidates, linePairs)
	}

	// for each line pair get the number of differences, if there are 0 total differences then we have a reflection
	for _, candidate := range reflectionCandidates {
		totalDifferences := 0

		for _, linePair := range candidate {
			totalDifferences += countDifferences(pattern, linePair)		
		}

		if totalDifferences == 0 {
			// return lines before the reflection point because it's a reflection
			noDifferenceLines = candidate[0][0] + 1 // because the problem is 1-indexed but our arrays are not
		}

		if totalDifferences == 1 {
			// this is part 2 which I'm not supposed to know but I spoiled myself on reddit, but there's only one difference means it's a "smudge"
			singleDifferenceLines = candidate[0][0] + 1 // because the problem is 1-indexed but our arrays are not
		}
	}

	log.SetPrefix(oldPrefix)
	return noDifferenceLines, singleDifferenceLines

}

func main() {
	log.SetPrefix(color.Green + "[ # 13 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")

	totalHorizontal := 0
	totalVertical := 0

	totalHorizontal2 := 0
	totalVertical2 := 0

	for _, pattern := range input {

		noDif, singleDif := processPattern(pattern)
		totalHorizontal += noDif
		totalHorizontal2 += singleDif

		transposedPattern := strutil.Transpose(pattern)

		noDif, singleDif = processPattern(transposedPattern)
		totalVertical += noDif
		totalVertical2 += singleDif
	}

	total := (100 * totalHorizontal) + totalVertical
	total2 := (100 * totalHorizontal2) + totalVertical2

	log.Println("Total:", total)
	log.Println("Total2:", total2)
}
