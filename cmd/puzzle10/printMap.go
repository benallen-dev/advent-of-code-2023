package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// Doesn't work properly for large inputs, but I solved the puzzle and I would like to move on
func printMap(input []string, visitedTiles map[Location]bool, insidePoints map[Location]bool) {

	// Print the map
	for linenum, line := range input {

		thisLine := ""

		for pointnum, point := range line {
			currentLocation := Location{linenum, pointnum}

			if visitedTiles[currentLocation] {
				thisLine += color.Blue + string(point) + color.Reset
			} else if insidePoints[currentLocation] {
				thisLine += color.Red + string(point) + color.Reset
			} else {
				thisLine += color.Gray + string(point) + color.Reset
			}
		}

		log.Println(thisLine)
	}
}
