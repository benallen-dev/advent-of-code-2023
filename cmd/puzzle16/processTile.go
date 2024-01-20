package main

import (
	"log"
)

// If we've come from the flat side of a splitter before we don't need to re-add the pointy ends to the queue
var visitedSplitters = map[Position]bool{}

func processTile(queueItem QueueItem) {
	position := queueItem.position
	previousPosition := queueItem.previousPosition

	tile := tileMap.GetTile(position)

	// Tile is energized
	energiseTile(position)

	// From which direction did we come?
	origin := ""
	if previousPosition[0] == position[0] {
		if previousPosition[1] < position[1] {
			origin = "left"
		} else {
			origin = "right"
		}
	} else {
		if previousPosition[0] < position[0] {
			origin = "top"
		} else {
			origin = "bottom"
		}
	}

	switch tile {
	case '.':
		// Beam goes straight
		nextPosition := position.Continue(previousPosition)
		queuePosition(nextPosition, position)
	case '-':
		if origin == "left" || origin == "right" {
			nextPosition := position.Continue(previousPosition)
			queuePosition(nextPosition, position)
		} else { // if coming from top or bottom, beam goes left AND right
			// Check if we've visited this splitter before
			visited := visitedSplitters[position]
			if !visited {
				queueRight(position)
				queueLeft(position)
				visitedSplitters[position] = true
			}

		}
	case '|':
		if origin == "top" || origin == "bottom" {
			nextPosition := position.Continue(previousPosition)
			queuePosition(nextPosition, position)
		} else { // if coming from left or right, beam goes up AND down
			// Check if we've visited this splitter before
			visited := visitedSplitters[position]
			if !visited {
				queueUp(position)
				queueDown(position)
				visitedSplitters[position] = true
			}

		}
	case '/':
		switch origin {
		case "top":
			queueLeft(position)
		case "left":
			queueUp(position)
		case "right":
			queueDown(position)
		case "bottom":
			queueRight(position)
		}
	case '\\': // Double backslash because otherwise it's an escape character
		switch origin {
		case "top":
			queueRight(position)
		case "left":
			queueDown(position)
		case "right":
			queueUp(position)
		case "bottom":
			queueLeft(position)
		}
	default:
		log.Fatal("Invalid tile, get rekt noobs ", string(tile))
	}
}
