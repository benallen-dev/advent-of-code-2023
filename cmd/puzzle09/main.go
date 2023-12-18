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

func predictPrevious (elements []int) int {
	// Determine intervals until they're all 0
	var intervals [][]int

	// Calculate the first set of intervals
	intervals = append(intervals, elements)

	// While the last set of intervals is not all 0
	for !allZero(intervals[len(intervals) - 1]) {
		// Calculate the next set of intervals
		intervals = append(intervals, calculateIntervals(intervals[len(intervals) - 1]))
	}


	

	for i := len(intervals) - 1; i >= 0; i-- {
		// If all zeros, add a 0 to the end
		if allZero(intervals[i]) {
			intervals[i] = append([]int{0}, intervals[i]...)
		} else {
			// Get the first element of the i'th array
			firstElement := intervals[i][0]
			// Get the first element of the i+1'th array
			nextFirstElement := intervals[i + 1][0]

			// prediction + nextFirstElement = firstElement
			intervals[i] = append([]int{firstElement - nextFirstElement}, intervals[i]...)
		}
	}

	// The first element of the first array is the answer
	return intervals[0][0]
}

func main() {
	log.SetPrefix(color.Green + "[ # 09 ] " + color.Reset)
	log.SetFlags(0)

	histories := readInput("input.txt")
	runningTotalNext := 0
	runningTotalPrevious := 0

	for _, history := range histories {
		runningTotalNext += predictNext(history)
		runningTotalPrevious += predictPrevious(history)
	}

	log.Println("Part 1:", runningTotalNext)
	log.Println("Part 2:", runningTotalPrevious)
}

