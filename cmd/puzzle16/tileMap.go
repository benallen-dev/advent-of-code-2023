package main

import (
	"fmt"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

type TileMap [][]rune

func (t TileMap) GetTile(position Position) rune {
	return t[position[0]][position[1]]
}

func (t TileMap) String() string {
	var str string

	str += "TileMap:\n"

	for _, row := range t {
		for _, tile := range row {
			str += string(tile)
		}
		str += "\n"
	}
	return str
}

func (t TileMap) PrintWithEnergy(energizedTiles [][]int) string {
	var str string

	str += "TileMap:\n"

	// TODO: Fix me
	for i, row := range t {
		for j, tile := range row {
			energy := energizedTiles[i][j]
			if energy > 0 {
				str += color.Yellow + fmt.Sprint(energy) + color.Reset
			} else {
				str += string(tile)
			}
		}
		str += "\n"
	}
	return str
}
