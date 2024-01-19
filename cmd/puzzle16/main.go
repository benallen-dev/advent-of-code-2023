package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// ait ait ait so, light beam starts at 0,0

var processQueue [][2]int
var tileMap TileMap
var energizedTiles []Position

func processTile(position Position, previousPosition Position) {
	log.Println("Processing tile", position)
	tile := tileMap.GetTile(position)

	// Tile is energized
	energizedTiles = append(energizedTiles, position)

	switch tile {
	case '.':
		// Beam goes straight
		nextPosition := position.Diff(previousPosition)

		// Add the next tile to the queue
		processQueue = append(processQueue, nextPosition)
	default:
		log.Fatal("Invalid tile, get rekt noobs ", string(tile))
	}

	// Set this tile as 'energized'
	tileMap[position[0]][position[1]] = '#'
}

func main() {
	log.SetPrefix(color.Green + "[ # 16 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("example.txt")
	for i, line := range input {
		// Create a new row
		tileMap = append(tileMap, make([]rune, len(line)))

		// Get rekt j-lovers
		for ii, char := range line {
			tileMap[i][ii] = char
		}
	}

	processQueue = append(processQueue, [2]int{0, 0}) // Let's gooo

	for len(processQueue) > 0 {
		nextItem := processQueue[0] // Get the first item
		processQueue = processQueue[1:] // Remove it from the queue

		processTile(Position(nextItem), Position([2]int{0, -1}))
		// Add to the queue
		// heck yeah the concepts works
	}

	log.Println("Done")
}


