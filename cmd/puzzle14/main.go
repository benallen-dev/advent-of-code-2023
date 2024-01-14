package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func calculateLoad(foo []string) (load int) {
	
	for i, line := range foo {
		rockLoad := len(foo) - i
		// count number of O's in line
		for _, char := range line {
			if char == 'O' {
				load += rockLoad
			}
		}
	}

	return load
}

func main() {
	log.SetPrefix(color.Green + "[ # 14 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")

	for _, line := range input {
		log.Println(line)
	}

	shifted := []string{
		"OOOO.#.O..",
		"OO..#....#",
		"OO..O##..O",
		"O..#.OO...",
		"........#.",
		"..#....#.#",
		"..O..#.O.O",
		"..O.......",
		"#....###..",
		"#....#....",
	}

	log.Println("Part 1: ", calculateLoad(shifted))
}
