package main

import (
	"log"
	"regexp"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// Location is a 2D coordinate, with the first element being the line number, and the second being the column number
type Location [2]int

/**
 * The pipes are arranged in a two-dimensional grid of tiles:
 *
 * | is a vertical pipe connecting north and south.
 * - is a horizontal pipe connecting east and west.
 * L is a 90-degree bend connecting north and east.
 * J is a 90-degree bend connecting north and west.
 * 7 is a 90-degree bend connecting south and west.
 * F is a 90-degree bend connecting south and east.
 * . is ground; there is no pipe in this tile.
 * S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
 */

func main() {
	log.SetPrefix(color.Green + "[ # 10 ] " + color.Reset)
	log.SetFlags(0)

	// input := readInput("exampleInput01.txt")
	// input := readInput("exampleInputPart2-3.txt")
	input := readInput("input.txt")

	// Find the start
	start, err := findStart(input)
	if err != nil {
		log.Fatal(err)
	}

	// previousTiles is the start tile in both cases at the start
	previousTiles := []Location{
		start,
		start,
	}
	// Get the two tiles that connect to the start
	tilesToVisit, err := getStartTiles(input, start)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1: How many tiles in the loop
	visitedTiles := map[Location]bool{}
	distance := 0
	visitedTiles[start] = true

	for tilesToVisit[0] != tilesToVisit[1] {
		// Update max distance
		distance++

		// Save distances in map
		visitedTiles[tilesToVisit[0]] = true
		visitedTiles[tilesToVisit[1]] = true

		// Visit both tiles, get next
		foo, err := getNextTile(input, tilesToVisit[0], previousTiles[0])
		if err != nil {
			log.Fatal(err)
		}
		bar, err := getNextTile(input, tilesToVisit[1], previousTiles[1])
		if err != nil {
			log.Fatal(err)
		}

		// Store tilesToVisit in previousTiles
		previousTiles = tilesToVisit // Lord I hope this isn't a reference lol
		tilesToVisit = []Location{
			foo,
			bar,
		}
	}

	// Save the last tile
	visitedTiles[tilesToVisit[0]] = true

	// At this point we have a map of all visited (loop) tiles
	// The distance variable holds the max distance - 1 as the loop was broken before the last iteration

	log.Println("Max distance is", distance+1)

	// Part 2: How many tiles enclosed by the loop
	// Strategy:
	// For each point
	//	- traverse left
	//	-	if odd tiles are part of the loopt, the point is contained inside it

	// a | means we're crossing the line
	// so does F-?J
	// so does L-?7
	edgeRegex := regexp.MustCompile(`(\||F-*J|L-*7)`)
	insidePoints := map[Location]bool{}
	insideCount := 0

	for linenum, line := range input {

		for pointnum := range line {

			currentLocation := Location{linenum, pointnum}
			currentIsVisited := visitedTiles[currentLocation]

			if currentIsVisited {
				continue
			}

			legitIntersections := [][]int{}

			intersections := edgeRegex.FindAllStringIndex(line[0:pointnum], -1)

			//			log.Println("intersections", intersections)

			for _, intersection := range intersections {
				if visitedTiles[Location{linenum, intersection[0]}] {
					// this intersection is legit
					legitIntersections = append(legitIntersections, intersection)
				}
			}

			//			log.Println("legitIntersections", legitIntersections)

			if len(legitIntersections)%2 == 1 {
				insidePoints[currentLocation] = true
				insideCount++
			}
		}
	}


	log.Println("There are", insideCount, "points inside the loop")
	
	// printMap(input, visitedTiles, insidePoints)
}
