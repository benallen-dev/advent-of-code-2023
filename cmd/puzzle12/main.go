package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetPrefix(color.Green + "[ # 12 ] " + color.Reset)
	log.SetFlags(0)

	springGroups := readInput("input.txt")

	totalArrangements := 0
	for _, springGroup := range springGroups {
		totalArrangements += springGroup.Arrangements()
	}

	log.Printf("Total arrangements part 1: %d", totalArrangements)
	
	// Now for the painful bit where my ineficcieny is punished
	unfoldedSpringGroups := unfoldRecords(springGroups)

	// LOL doing this filled up 32GB of RAM in like 20 seconds
	totalArrangementsPart2 := 0
	for _, unfoldedSpringGroup := range unfoldedSpringGroups {
		totalArrangementsPart2 += unfoldedSpringGroup.Arrangements()
	}

	log.Printf("Total arrangements part 2: %d", totalArrangementsPart2)
}
