package main

import (
	"fmt"
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var tileMap TileMap
var energizedTiles [][]int

func energiseTile(position Position) {
	i := position[0]
	j := position[1]

	energizedTiles[i][j] += 1
}

// This so much spaghetti in just one function, relies on the global tileMap
func isOutOfBounds(p Position) bool {
	return p[0] < 0 || p[1] < 0 || p[0] >= len(tileMap) || p[1] >= len(tileMap[0])
}

func countEnergizedTiles() int {
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
	for i, line := range tileMap {
		energizedTiles = append(energizedTiles, make([]int, len(line)))

		for ii := range line {
			energizedTiles[i][ii] = 0
		}
	}

	positionQueue.initWith(position, previousPosition)

	for !positionQueue.isEmpty() {
		queueItem := positionQueue.pop()
		processTile(queueItem)
	}

	count := countEnergizedTiles()
	fmt.Println("Energized tiles:", count)
	return count
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

	log.Println("Done")
}
