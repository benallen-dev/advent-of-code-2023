package main

import (
	"errors"
)

// getTile returns the rune at the given location
//
// Example:
//
//	input := []string{"F-7", "|.|", "L-J"}
//	getTile(input, Location{0, 0}) // 'F', nil
func getTile(input []string, location Location) (tile rune, err error) {
	if location[0] < 0 || location[0] >= len(input) {
		return ' ', errors.New("Out of bounds")
	}

	if location[1] < 0 || location[1] >= len(input[location[0]]) {
		return ' ', errors.New("Out of bounds")
	}

	return rune(input[location[0]][location[1]]), nil
}
