package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func calculateIntervals (elements []int) []int {
	result := []int{}

	for i := 1; i < len(elements); i++ {
		result = append(result, elements[i] - elements[i - 1])
	}

	return result
}

func allZero (elements []int) bool {
	for _, element := range elements {
		if element != 0 {
			return false
		}
	}

	return true
}

func predictNext (elements []int) int {
	// Implement the algorithm explained in the puzzle

	// Determine intervals until they're all 0
	var intervals [][]int

	// Calculate the first set of intervals
	intervals = append(intervals, elements)

	// While the last set of intervals is not all 0
	for !allZero(intervals[len(intervals) - 1]) {
		// Calculate the next set of intervals
		intervals = append(intervals, calculateIntervals(intervals[len(intervals) - 1]))
	}

	// We can now start adding to the ends of arrays and whatnot

	for i := len(intervals) - 1; i >= 0; i-- {
		// If all zeros, add a 0 to the end
		if allZero(intervals[i]) {
			intervals[i] = append(intervals[i], 0)
		} else {
			// Get the last element of the i'th array
			lastElement := intervals[i][len(intervals[i]) - 1]
			//Get the last element of the i+1'th array
			nextLastElement := intervals[i + 1][len(intervals[i + 1]) - 1]
			// In the example, this is 3 and 0
			// add these together and append to the i'th array
			intervals[i] = append(intervals[i], lastElement + nextLastElement)
		}
	}

	// The last element of the first array is the answer
	return intervals[0][len(intervals[0]) - 1]
	
}

func main() {
	log.SetPrefix(color.Green + "[ # 09 ] " + color.Reset)
	log.SetFlags(0)

	histories := readInput("input.txt")
	runningTotal := 0

	for _, history := range histories {
		runningTotal += predictNext(history)
	}

	log.Println("Part 01:", runningTotal)


	// Part 2 is gonna be brutal if they explain step by step
	// how to solve part 1



}

