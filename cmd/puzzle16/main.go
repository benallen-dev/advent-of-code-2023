package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var tileMap TileMap
var energizedTiles [][]int
var visitedSplitters = map[Position]bool{}

func energiseTile(position Position) {
	i := position[0]
	j := position[1]

	energizedTiles[i][j] += 1
}

// This so much spaghetti in just one function, relies on the global tileMap
func isOutOfBounds(p Position) bool {
	return p[0] < 0 || p[1] < 0 || p[0] >= len(tileMap) || p[1] >= len(tileMap[0])
}

func countEnergizedTiles(energizedTiles [][]int) int {
	count := 0
	for _, row := range energizedTiles {
		for _, tile := range row {
			if tile > 0 {
				count++
			}
		}
	}
	return count
}

func testConfiguration(tileMap TileMap, position Position, previousPosition Position) int {
	// reset energizedTiles
	for i, line := range tileMap {
		energizedTiles = append(energizedTiles, make([]int, len(line)))

		for ii := range line {
			energizedTiles[i][ii] = 0
		}
	}

	// reset visitedSplitters
	visitedSplitters = map[Position]bool{}

	positionQueue.initWith(position, previousPosition)

	for !positionQueue.isEmpty() {
		queueItem := positionQueue.pop()
		processTile(queueItem)
	}

	return countEnergizedTiles(energizedTiles)
}

// This program is at total mess lol
func main() {
	log.SetPrefix(color.Green + "[ # 16 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")
	for i, line := range input {
		// initialise the maps
		tileMap = append(tileMap, make([]rune, len(line)))

		// Get rekt j-lovers we're using ii
		for ii, char := range line {
			tileMap[i][ii] = char
		}
	}

	part01 := testConfiguration(tileMap, Position{0, 0}, Position{0, -1})
	log.Println("Part 01:", part01)

	startingPositions := []QueueItem{}

	// Part 02
	// Generate a list of starting positions
	// First all the top and bottom positions
	colHeight := len(tileMap)

	for i := 0; i < len(tileMap[0]); i++ {
		startingPositions = append(startingPositions, QueueItem{Position{i, 0}, Position{i, -1}})
		startingPositions = append(startingPositions, QueueItem{Position{i, colHeight - 1}, Position{i, colHeight}})
	}

	// Then all the left and right positions
	rowWidth := len(tileMap[0])
	for i := 0; i < len(tileMap); i++ {
		startingPositions = append(startingPositions, QueueItem{Position{0, i}, Position{-1, i}})
		startingPositions = append(startingPositions, QueueItem{Position{rowWidth - 1, i}, Position{rowWidth, i}})
	}

	// Test all the starting positions and get the highest energized tile count
	highestEnergizedTileCount := 0
	for _, startingPosition := range startingPositions {
		energizedTileCount := testConfiguration(tileMap, startingPosition.position, startingPosition.previousPosition)
		if energizedTileCount > highestEnergizedTileCount {
			highestEnergizedTileCount = energizedTileCount
		}
	}

	log.Println("Part 02:", highestEnergizedTileCount)

}
