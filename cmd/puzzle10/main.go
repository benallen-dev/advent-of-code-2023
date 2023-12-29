package main

import (
	"errors"
	"log"

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

 // getNextTile takes the input, the current location, and the previous location, and returns the next location
 // It also returns an error if the current tile is invalid or the direction is invalid
func getNextTile(input []string, location Location, prevLocation Location) (next Location, err error) {

	tile, err := getTile(input, location)
	if err != nil {
		return Location{}, err
	}

	// I'm sure there's a more efficient way of doing this but "verbose and works" is better than "clever and broken"

	// Check if we're coming from NESW
	if location[0] > prevLocation[0] { // Coming from North
		switch tile {
		case '|': // return tile to the south of us
			return Location{location[0] + 1, location[1]}, nil
		case 'L':
			return Location{location[0], location[1] + 1}, nil
		case 'J':
			return Location{location[0], location[1] - 1}, nil
		default:
			return Location{}, errors.New("Invalid tile from North: " + string(tile))
		}
	} else if location[0] < prevLocation[0] { // South
		switch tile {
		case '|':
			return Location{location[0] - 1, location[1]}, nil
		case 'F':
			return Location{location[0], location[1] + 1}, nil
		case '7':
			return Location{location[0], location[1] - 1}, nil
		default:
			return Location{}, errors.New("Invalid tile from South: " + string(tile))
		}
	} else if location[1] < prevLocation[1] { // East
		switch tile {
		case '-':
			return Location{location[0], location[1] - 1}, nil
		case 'L':
			return Location{location[0] - 1, location[1]}, nil
		case 'F':
			return Location{location[0] + 1, location[1]}, nil
		default:
			return Location{}, errors.New("Invalid tile from East: " + string(tile))
		}
	} else if location[1] > prevLocation[1] { // West
		switch tile {
		case '-':
			return Location{location[0], location[1] + 1}, nil
		case 'J':
			return Location{location[0] - 1, location[1]}, nil
		case '7':
			return Location{location[0] + 1, location[1]}, nil
		default:
			return Location{}, errors.New("Invalid tile from West: " + string(tile))
		}
	} else {
		log.Println("Error getting next tile from", location, "coming from", prevLocation)
		return Location{}, errors.New("Not coming from NESW")
	}
}

func main() {
	log.SetPrefix(color.Green + "[ # 10 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("exampleInputPart2.txt")
	// input := readInput("input.txt")

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

	// At this point we have a map of all visited (loop) tiles
	// The distance variable holds the max distance - 1 as the loop was broken before the last iteration

	log.Println("Max distance is", distance + 1)



	// Part 2: How many tiles enclosed by the loop
	// Strategy:
	// For each point
	//	- traverse left
	//	-	if odd tiles are part of the loopt, the point is contained inside it


}
