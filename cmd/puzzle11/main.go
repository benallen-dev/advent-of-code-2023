package main

import (
	"log"
	"math"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

type Location [2]int

func findDistanceInGiantUniverse(a Location, b Location, emptyLines []int, emptyColumns []int) (distance int) {

	xDistance := math.Abs(float64(a[0] - b[0]))
	yDistance := math.Abs(float64(a[1] - b[1]))
	
	// We need to add 999_999 to the distance because empty lines represent a distance of 1_000_000, not 1_000_001 like I did the first time
	for _, line := range emptyLines {
		// This if-statement is afwul to read
		// TLDR if line is inbetween a and b, add 999_999 to the distance
		if (a[0] < b[0] && a[0] < line && b[0] > line) || (a[0] > b[0] && a[0] > line && b[0] < line) {
			yDistance += 999_999
		}
	}

	// Same for columns
	// This is a bit easier to read and understand I think but it boils down to the same thing
	for _, column := range emptyColumns {
		if a[1] < b[1] && a[1] < column && b[1] > column {
			yDistance += 999_999
		} else if a[1] > b[1] && a[1] > column && b[1] < column {
			yDistance += 999_999
		}
	}

	return int(xDistance + yDistance)
}

func main() {
	log.SetPrefix(color.Green + "[ # 11 ] " + color.Reset)
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

	// Now we need to determine distances, except empty lines are now x1_000_000

	// from here on we can't use the "universe" variable, we need to use math to compute the distances
	emptyLines := findEmptyLines(input)
	emptyColumns := findEmptyColumns(input)

	// We can just reuse galaxies as we already printed out the total distance for part 1
	// This is now a non-expanded universe, just for the record
	galaxies = findGalaxies(input)
	// We can reuse totalDistance as well
	totalDistance = 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			if i < j {
				distance := findDistanceInGiantUniverse(galaxies[i], galaxies[j], emptyLines, emptyColumns)
				totalDistance += distance
			}
		}
	}

	log.Println("Total distance in giant universe:", totalDistance)
}
