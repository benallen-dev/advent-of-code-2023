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
	// LOL doing this filled up 32GB of RAM in like 20 seconds
	
	// This is just straight-up not going to work because the numbers are too big.
	// The first line alone contains 49 "?" characters, which means 2^49 possibilities.
	// Even if you stop generating the tree once the groups no longer match that's still
	// an astronomical number.

	// William Y. Feng says in this video that he solved it using dynamic programming.
	// https://www.youtube.com/watch?v=veJvlIMjv94
	// Which is cool because I've heard of dynamic programming but I have yet to actually use it.
	// I think I've bashed my head against the wall enough to call this a learning opportunity
	// and try to implement his solution (originally in Python) in Go.

	// This is the code I initially used to try and solve the problem, left here for posterity.
	// Feel free to laugh at me for being so naive, I learned a lot!
	// unfoldedSpringGroups := unfoldRecords(springGroups)
	// totalArrangementsPart2 := 0
	//
	// for idx, unfoldedSpringGroup := range unfoldedSpringGroups {
	// 	log.Printf("Unfolded spring group: %s", unfoldedSpringGroup)
	//
	// 	log.Printf("Generating possibility tree...")
	// arrangements := unfoldedSpringGroup.Arrangements()
	// log.Printf("Line %d of %d: %d", idx, len(unfoldedSpringGroups), arrangements)
	//
	// totalArrangementsPart2 += arrangements
	//
	// }

	// log.Printf("Total arrangements part 2: %d", totalArrangementsPart2)
}
