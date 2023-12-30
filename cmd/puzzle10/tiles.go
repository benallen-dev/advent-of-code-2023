package main

import (
	"log"
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
		return ' ', errors.New("GetTile :: Out of bounds")
	}

	if location[1] < 0 || location[1] >= len(input[location[0]]) {
		return ' ', errors.New("GetTile :: Out of bounds")
	}

	return rune(input[location[0]][location[1]]), nil
}

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
