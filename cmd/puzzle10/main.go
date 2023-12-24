package main

import (
	"errors"
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

const (
	North = iota
	South
	East
	West
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

// So, basic premise: there exists a loop of pipes, and given a starting position, we want to find the point halfway around the loop.

// Approach 1:
// Find S
// Pass to the following function
// given an element X
// if X already has a number, we're going back around the loop, exit.
//   ( it should have value max - prev )
// else
//   Assign a value count + 1 to X
// for each element around X, find elements that connect to X
//
//

func findStart(input []string) (location Location, err error) {
	for lineIndex, line := range input {
		for colIndex, char := range line {
			if char == 'S' {
				return Location{lineIndex, colIndex}, nil
			}
		}
	}

	return Location{-1, -1}, errors.New("Could not find start")
}

func getTile(input []string, location Location) (tile rune, err error) {
	if location[0] < 0 || location[0] >= len(input) {
		return ' ', errors.New("Out of bounds")
	}

	if location[1] < 0 || location[1] >= len(input[location[0]]) {
		return ' ', errors.New("Out of bounds")
	}

	return rune(input[location[0]][location[1]]), nil
}

func main() {
	log.SetPrefix(color.Green + "[ # 10 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("exampleInput01.txt")

	// Find the start
	start, err := findStart(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start is at", start)

	// Find elements around S that are pipes that connect to S
	// Luckily, we can just check the 4 cardinal directions
	north := [2]int{start[0] - 1, start[1]}
	south := [2]int{start[0] + 1, start[1]}
	// east := [2]int{start[0], start[1] + 1}
	// west := [2]int{start[0], start[1] - 1}

	// First, North
	northTile, err := getTile(input, north)
	if err != nil {
		log.Println(err)
	} else {
		switch northTile {
		case '|':
			log.Println("North exits north")
		case 'F':
			log.Println("North exists east")
		case '7':
			log.Println("North exists west")
		default:
			log.Println("North is not connected")
		}
	}

	// Second, South
	
	southTile, err := getTile(input, south)
	if err != nil {
		log.Println(err)
	} else {
		switch southTile {
		case '|':
			log.Println("South exits south")
		case 'J':
			log.Println("South exists west")
		case 'L':
			log.Println("South exists east")
		default:
			log.Println("South is not connected")
		}
	}

}
