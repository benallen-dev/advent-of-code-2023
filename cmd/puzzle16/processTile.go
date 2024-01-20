package main

import (
	"log"
)

func processTile(queueItem QueueItem) {
	position := queueItem.position
	previousPosition := queueItem.previousPosition

	log.Println("Processing tile", position, "from", previousPosition)

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
			queueLeft(position)
			queueRight(position)
		}
	case '|':
		if origin == "top" || origin == "bottom" {
			nextPosition := position.Continue(previousPosition)
			queuePosition(nextPosition, position)
		} else { // if coming from left or right, beam goes up AND down
			queueUp(position)
			queueDown(position)
		}
	case '/':
		// If coming from top, beam goes left
		// If coming from left, beam goes top
		// If coming from right, beam goes down
		// If coming from bottom, beam goes right
	case '\\': // Double backslash because otherwise it's an escape character
		// If coming from top, beam goes right
		// If coming from left, beam goes down
		// If coming from right, beam goes top
		// If coming from bottom, beam goes left
	default:
		log.Fatal("Invalid tile, get rekt noobs ", string(tile))
	}
}
