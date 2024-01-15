package main

import (
	"log"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
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

func shiftNorth(pattern []string) (shifted []string) {

	// Create strings from the columns of pattern
	columnStrings := strutil.Transpose(pattern)

	// Split the strings on the '#' character
	for _, column := range columnStrings {
		groups := strings.Split(column, "#")
		newGroups := []string{}
		
		for _, group := range groups {

			// Count O's
			oCount := 0
			for _, char := range group {
				if char == 'O' {
					oCount++
				}
			}

			// Bunch em up at the beginning
			newString := strings.Repeat("O", oCount) + strings.Repeat(".", len(group)-oCount)
			newGroups = append(newGroups, newString)
		}

		// Join the groups back together with '#' characters
		newColumn := strings.Join(newGroups, "#")
		shifted = append(shifted, newColumn)
	}

	// transpose the shifted pattern
	shifted = strutil.Transpose(shifted)

	return shifted
}

func main() {
	log.SetPrefix(color.Green + "[ # 14 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")

	shifted := shiftNorth(input)

	// exampleShifted := []string{
	// 	"OOOO.#.O..",
	// 	"OO..#....#",
	// 	"OO..O##..O",
	// 	"O..#.OO...",
	// 	"........#.",
	// 	"..#....#.#",
	// 	"..O..#.O.O",
	// 	"..O.......",
	// 	"#....###..",
	// 	"#....#....",
	// }

	log.Println("Part 1: ", calculateLoad(shifted))
}
