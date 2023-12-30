package main

import (
	"log"
	"math"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

type Location [2]int

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

func main() {
	log.SetPrefix(color.Green + "[ # 01 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")
	universe := expandUniverse(input)
	galaxies := findGalaxies(universe)

	totalDistance := 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			if i < j {
				distance := findDistance(galaxies[i], galaxies[j])
				totalDistance += distance
			}
		}
	}

	log.Println("Total distance:", totalDistance)
}
