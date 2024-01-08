package main

import (
	"fmt"
)


// getGroupNumbers takes a SpringGroup.strings and returns a slice of integers that represent contiguous groups of damaged springs
// It will not be happy if you pass it a string containing ? characters
func getGroupNumbers(springs string) ([]int, error) {
	groups := []int{}
	currentGroup := 0
	
	var previousSpring rune

	for idx, spring := range springs {
		if spring == '#' {
			currentGroup++
		} else if spring == '.' {
			if previousSpring == '#' {
				groups = append(groups, currentGroup)
				currentGroup = 0
			}
		} else {
			return nil, fmt.Errorf("getGroupNumbers: invalid spring character: %c", spring)
		}

		if idx == len(springs) - 1 && currentGroup > 0 {
			groups = append(groups, currentGroup)
		}
		
		previousSpring = spring
	}

	return groups, nil
}

