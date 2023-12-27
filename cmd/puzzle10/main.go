package main

import (
	"errors"
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// Location is a 2D coordinate, with the first element being the line number, and the second being the column number
type Location [2]int

type Pipe struct {
	Location Location
	Input    int
	Output   int
}

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

	input := readInput("exampleInput01.txt")

	// Find the start
	start, err := findStart(input)
	if err != nil {
		log.Fatal(err)
	}

	// Get the two tiles that connect to the start
	startTiles, err := getStartTiles(input, start)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Start is at", start)
	log.Println("Next tiles are", startTiles)
	
	// tileDistances := map[Location]int{}

}
