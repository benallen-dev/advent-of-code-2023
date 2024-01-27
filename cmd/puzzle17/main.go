package main

import (
	"log"
	"math"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func hasUnvisited() bool {
	log.Panic("Not implemented: getLowestUnvisited")
	return false
}

func getLowestUnvisited() int {
	log.Panic("Not implemented: getLowestUnvisited")
	return -1
}

func main() {
	log.SetPrefix(color.Green + "[ # 17 ] " + color.Reset)
	log.SetFlags(0)

	// The supplied input can be interpreted as a 2D grid of node weights, where
	// we want to find a path with the lowest weight. Apart from the limitation
	// on the number of consecutive steps in a single direction, this is a
	// shortest path problem ,so time to break out Dijkstra's algorithm.
	adjs := readInput("example.txt")

	for _, adj := range adjs {
		for _, edge := range adj {
		log.Println(edge)
		}
	}

	// Fill dists with near-as-makes-no-difference infinity
	// Fill seen with false
	dists := make([]int, len(adjs))
	seen  := make([]bool, len(adjs))

	for i := range adjs {
		dists[i] = math.MaxInt
		seen[i] = false
	}

	source := 0
	target := len(adjs) - 1
	dists[source] = 0

	log.Printf("Source: %d, Target: %d", source, target)


	// Let's test the minheap for a sec


	// Build a minheap of all the distances
	minheap := NewMinheap()
	for i := range adjs {
		minheap.Insert(i, math.MaxInt)
	}

	// Update the source node to have a distance of 0
	minheap.Update(source, 0)



// 	// Time to algorithm!
// 	for hasUnvisited() == true {
// 		lo := getLowestUnvisited()
// 		seen[lo] = true

// 		edges := adjs[lo]

// 		// adjust the adj to take the 'no more than 3 steps in one direction'
// 		// rule into account
// 		for i, edge := range edges {
// 			if seen[i] {
// 				continue
// 			}

// 			dist := dists[lo] + edge.weight
// 			if dist < dists[i] {
// 				dists[edge.from] = dist
// 			}
// 		}
// 	}
}
