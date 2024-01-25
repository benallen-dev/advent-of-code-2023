package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetPrefix(color.Green + "[ # 17 ] " + color.Reset)
	log.SetFlags(0)


	// The supplied input can be interpreted as a 2D grid of node weights, where
	// we want to find a path with the lowest weight. Apart from the limitation
	// on the number of consecutive steps in a single direction, this is a
	// shortest path problem ,so time to break out Dijkstra's algorithm.
	adjs := readInput("example.txt")

	for i, adj := range adjs {
		log.Printf("%2d : %v", i, adj)
	}


}
