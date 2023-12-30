package main

import (
	"math"
)

func findDistance(a, b Location) int {
	xDistance := math.Abs(float64(a[0] - b[0]))
	yDistance := math.Abs(float64(a[1] - b[1]))

	return int(xDistance + yDistance)
}

func findGalaxies(universe []string) (galaxies map[int]Location) {
	galaxies = make(map[int]Location)
	index := 0

	for i, line := range universe {
		for j, char := range line {
			if char == '#' {
				galaxies[index] = Location{i, j}
				index++
			}
		}
	}

	return galaxies
}
