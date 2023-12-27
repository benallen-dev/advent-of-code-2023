package main

import (
	"errors"
)

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

func getStartTiles(input []string, start Location) (tiles []Location, err error) {
	// Find elements around S that are pipes that connect to S
	// Luckily, we can just check the 4 cardinal directions
	north := [2]int{start[0] - 1, start[1]}
	south := [2]int{start[0] + 1, start[1]}
	east := [2]int{start[0], start[1] + 1}
	west := [2]int{start[0], start[1] - 1}

	startTiles := []Location{}

	// For all of these append operations, it can theoretically go out of
	// bounds but we're making the admittedly enormous assumption that the input is
	// well-formatted. It's not user input after all.

	// First, North
	northTile, err := getTile(input, north)
	if err != nil {
		return nil, err
	} else {
		if northTile == '|' || northTile == 'F' || northTile == '7' {
			startTiles = append(startTiles, north)
		}
	}

	// Second, South
	southTile, err := getTile(input, south)
	if err != nil {
		return nil, err
	} else {
		if southTile == '|' || southTile == 'J' || southTile == 'L' {
			startTiles = append(startTiles, south)
		}
	}

	// Third, East
	eastTile, err := getTile(input, east)
	if err != nil {
		return nil, err
	} else {
		if eastTile == '-' || eastTile == 'J' || eastTile == '7' {
			startTiles = append(startTiles, east)
		}
	}

	// Fourth, West
	westTile, err := getTile(input, west)
	if err != nil {
		return nil, err
	} else {
		if westTile == '-' || westTile == 'L' || westTile == 'F' {
			startTiles = append(startTiles, west)
		}
	}

	return startTiles, nil
}
