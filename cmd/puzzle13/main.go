package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

func findDoubleLines(pattern []string) (lines [][]int) {
	for i := 0; i < len(pattern)-1; i++ { // for each line
		if pattern[i] == pattern[i+1] {
			lines = append(lines, []int{i, i + 1})
		}
	}

	return lines
}

func generateLinePairs(pattern []string, doubleLines []int) (linePairs [][]int) {
	// Save the old log prefix
	// oldPrefix := log.Prefix()
	// log.SetPrefix(color.Green + "[ # 13 :: generateLinePairs ] " + color.Reset)
	
	// [2,3] with length 6 means we have
	// [2, 3]
	// [1, 4]
	// [0, 5]

	left := doubleLines[0]
	right := doubleLines[1]

	// log.Println()

	// log.Println("Pattern length:", len(pattern))

	// log.Println("Left:", left)
	// log.Println("Right:", right)

	distanceToEnd := len(pattern) -1 - right
	iterations := min(left, distanceToEnd)

	// log.Println("Distance to end:", distanceToEnd)
	// log.Println("Iterations:", iterations)
	// log.Println()

	for i := 0; i <= iterations; i++ {
		linePairs = append(linePairs, []int{left - i, right + i})
	}

	// Restore the old log prefix
	// log.SetPrefix(oldPrefix)

	// We now should have a set of lines that doesn't dip below 0 or go above the length of the pattern
	return linePairs
}

func countDifferences(pattern []string, linePair []int) (differences int) {
	// oldPrefix := log.Prefix()
	// log.SetPrefix(color.Purple + "[ # 13 :: countDifferences ] " + color.Reset)
	
	differences = 0 // Not sure if init is needed but let's be explicit

	a := pattern[linePair[0]]
	b := pattern[linePair[1]]

	// log.Println("Line pair:", linePair)
	// log.Println("A:", a)
	// log.Println("B:", b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}

	// log.SetPrefix(oldPrefix)
	return differences
}


func processPattern(pattern []string) (noDifferenceLines int, singleDifferenceLines int) {

	oldPrefix := log.Prefix()
	log.SetPrefix(color.Cyan + "[ # 13 :: processPattern ] " + color.Reset)

	// find all sets of double lines
	// for each line pair get the number of differences
	// if the number of differences is 0, then we have a reflection for those lines
	// if all line pairs have a reflection, then we have a reflection for the whole pattern and we can count it
	doubleLines := findDoubleLines(pattern)

	// log.Println("Pattern length:", len(pattern))
	// log.Println("Double lines:", doubleLines)

	// for _, line := range pattern {
	// 	log.Println(line)
	// }

	reflectionCandidates := [][][]int{}

	// for each set of double lines, generate a list of line pairs
	for _, doubleLine := range doubleLines {
		linePairs := generateLinePairs(pattern, doubleLine)
		// log.Println("Generated Line pairs for", doubleLine)
		reflectionCandidates = append(reflectionCandidates, linePairs)
	}

	for _, candidate := range reflectionCandidates {
		totalDifferences := 0

//		log.Println("[")
		for _, linePair := range candidate {
			differences := countDifferences(pattern, linePair)
			totalDifferences += differences
//			log.Println("  ", linePair, ":", differences)
		}
//		log.Println("]")

		if totalDifferences == 0 {
			// return lines before the reflection point because it's a reflection
			noDifferenceLines = candidate[0][0] + 1 // because the problem is 1-indexed but our arrays are not
		}

		if totalDifferences == 1 {
			// this is part 2 which I'm not supposed to know but I spoiled myself on reddit
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
		// log.Println()
		// log.Println("Pattern:", i)

// 		log.Println("--------------------")
// 		for _, line := range pattern {
// 			log.Println(line)
// 		}
// 		log.Println("--------------------")

		noDif, singleDif := processPattern(pattern)
		totalHorizontal += noDif
		totalHorizontal2 += singleDif

		transposedPattern := strutil.Transpose(pattern)
//		log.Printf("Transposed pattern %d", i)

		// log.Println("--------------------")
		// for _, line := range transposedPattern {
		// 	log.Println(line)
		// }
		// log.Println("--------------------")

		noDif, singleDif = processPattern(transposedPattern)
		totalVertical += noDif
		totalVertical2 += singleDif
	}

	total := (100 * totalHorizontal) + totalVertical
	total2 := (100 * totalHorizontal2) + totalVertical2

//	log.Println()
	log.Println("Total:", total)
	log.Println("Total2:", total2)
}
