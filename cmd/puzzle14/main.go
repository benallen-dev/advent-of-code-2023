package main

import (
	"log"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

var cycleCache = map[string][]string{}

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

func cycle(pattern []string) (cycled []string) {
	// Check cache	
	key := strings.Join(pattern, "")
	if cached, ok := cycleCache[key]; ok {
		return cached
	}

	cycled = shiftNorth(pattern)
	cycled = shiftWest(cycled)
	cycled = shiftSouth(cycled)
	cycled = shiftEast(cycled)

	// Cache the result
	cycleCache[key] = cycled

	return cycled
}

func mirrorHorizontal(pattern []string) (mirrored []string) {
	for _, line := range pattern {
		mirrored = append(mirrored, strutil.Reverse(line))
	}

	return mirrored
}

func printPattern(pattern []string) {
	for _, line := range pattern {
		log.Println(line)
	}
}

func main() {
	log.SetPrefix(color.Green + "[ # 14 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("example.txt")

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

	log.Println("Part 1: ", calculateLoad(shiftNorth(input)))
	log.Println()

	cycled := input

	for i := 0; i < 1_000_000_000; i++ {
		if i%100_000 == 0 && i != 0 {
			percentDone := float32(i) / float32(1_000_000_0)
			log.Println("Progress:", percentDone, "%")
		}

		cycled = cycle(cycled)
	}

	log.Println("Part 2: ", calculateLoad(cycled))

}
