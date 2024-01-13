package main

import (
	"fmt"
	"strings"
)

var (
	cache = make(map[string][]int)
)

// getGroupNumbers takes a SpringGroup.strings and returns a slice of integers that represent contiguous groups of damaged springs
// It will not be happy if you pass it a string containing ? characters
func getGroupNumbers(springs string) ([]int, error) {

	// Memoize, let's see if it helps
	if cache[springs] != nil {
		return cache[springs], nil
	}

	groups := []int{}
	currentGroup := 0
	
	var previousSpring rune

	if strings.Contains(springs, "?") {
		return []int{}, fmt.Errorf("getGroupNumbers: cannot handle string with '?' character: %s", springs)
	}

	for idx, spring := range springs {
		if spring == '#' {
			currentGroup++
		} else if spring == '.' {
			if previousSpring == '#' {
				groups = append(groups, currentGroup)
				currentGroup = 0
			}
		} else {
			return nil, fmt.Errorf("getGroupNumbers: invalid spring character: '%c' in %s, index %d", spring, springs, idx)
		}

		if idx == len(springs) - 1 && currentGroup > 0 {
			groups = append(groups, currentGroup)
		}
		
		previousSpring = spring
	}

	cache[springs] = groups

	return groups, nil
}

