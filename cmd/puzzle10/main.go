package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

/**
 * The pipes are arranged in a two-dimensional grid of tiles:
 *
 * | is a vertical pipe connecting north and south.
 * - is a horizontal pipe connecting east and west.
 * L is a 90-degree bend connecting north and east.
 * J is a 90-degree bend connecting north and west.
 * 7 is a 90-degree bend connecting south and west.
 * F is a 90-degree bend connecting south and east.
 * . is ground; there is no pipe in this tile.
 * S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
 */


 // So, basic premise: there exists a loop of pipes, and given a starting position, we want to find the point halfway around the loop.

 // Approach 1:
 // Find S
 // Pass to the following function
 // given an element X
 // if X already has a number, we're going back around the loop, exit.
 //   ( it should have value max - prev )
 // else
 //   Assign a value count + 1 to X
 // for each element around X, find elements that connect to X
 // 
 // 

type PipeTraverser struct {
	length int
	// Line,Column
	location [2]int
}

func main() {
	log.SetPrefix(color.Green + "[ # 10 ] " + color.Reset)
	log.SetFlags(0)

	foo := readInput("input.txt")
	for _, line := range foo {
		log.Println(line)
	}

	bar := PipeTraverser{
		length: 0,
		location: [2]int{0,0},
	}

	bar.location[1] = bar.location[1] + 1

}
