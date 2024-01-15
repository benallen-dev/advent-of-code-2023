package main

import (
	"strings"

	"crypto/sha1"

	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

var (
	shiftLeftCache = map[[20]byte]string{}
)

func shiftLeft(pattern []string) (shifted []string) {
	// Split the strings on the '#' character
	for _, line := range pattern {
		key := sha1.Sum([]byte(line))

		if cached, ok := shiftLeftCache[key]; ok {
			shifted = append(shifted, cached)
			continue
		}

		groups := strings.Split(line, "#")
		newGroups := []string{}
		
		for _, group := range groups {

			// Count O's
			oCount := 0
			for _, char := range group {
				if char == 'O' {
					oCount++
				}
			}

			// Bunch em up at the beginning
			newString := strings.Repeat("O", oCount) + strings.Repeat(".", len(group)-oCount)
			newGroups = append(newGroups, newString)
		}

		// Join the groups back together with '#' characters
		newColumn := strings.Join(newGroups, "#")
		shifted = append(shifted, newColumn)

		// Cache the result
		shiftLeftCache[key] = newColumn
	}

	return shifted
}

func shiftNorth(pattern []string) (shifted []string) {
	// Create strings from the columns of pattern
	columnStrings := strutil.Transpose(pattern)

	shifted = shiftLeft(columnStrings)
	// Transpose the shifted columns back into rows
	shifted = strutil.Transpose(shifted)
	
	return shifted
}

func shiftWest(pattern []string) (shifted []string) {
	return shiftLeft(pattern)
}

func shiftEast(pattern []string) (shifted []string) {
	shifted = mirrorHorizontal(pattern)

	shifted = shiftLeft(shifted)

	// Un-reverse the strings from the columns of pattern
	shifted = mirrorHorizontal(shifted)

	return shifted
}

func shiftSouth(pattern []string) (shifted []string) {
	// Transpose then reverse
	shifted = strutil.Transpose(pattern)
	shifted = mirrorHorizontal(shifted)

	// Shift left
	shifted = shiftLeft(shifted)

	// Un-reverse then un-transpose
	shifted = mirrorHorizontal(shifted)
	shifted = strutil.Transpose(shifted)

	return shifted
}

