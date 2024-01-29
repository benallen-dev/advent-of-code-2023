package main

import (
	"fmt"
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var largeInt int = 10000000000 // But still less than MaxInt to avoid overflow

func hasUnvisited(seen []bool, dists []int) bool {
	for i := range seen {
		if !seen[i] && dists[i] < 10000000000 {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	index := -1
	lowest := largeInt

	// O(N^2) so let's hope part 2 doesn't blow up the input size
	for i := range seen {
		if !seen[i] && dists[i] < lowest {
			log.Printf("Found new lowest: %d", dists[i])
			index = i
			lowest = dists[i]
		}
	}

	return index
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func printGrid(lines []string, path []int) {
	out := ""

	for i, line := range lines {
		for j, char := range line {
			index := (i * len(lines[0])) + j
			// if in path, add a color
			if contains(path, index) {
				out += fmt.Sprintf("%s%c%s", color.Red, char, color.Reset)
			} else {
				out += fmt.Sprintf("%c", char)
			}
		}
		out += "\n"
	}

	log.Printf("Grid:\n%s", out)
}

func isFourthMove(curr int, lineLength int, next int, prev []int) (isFourthStep bool) {
	previous := prev[curr]
	if previous == -1 {
		return false
	}
	compareMe := prev[previous]
	if compareMe == -1 {
		return false
	}

	// Based on which direction we're going, check prev[prev[prev[curr]]] - if it's in the same direction discard this move
	if next == curr+1 {
		// coming from Left
		coord := next % lineLength
		compareCoord := compareMe % lineLength

		log.Printf("coord: %d, compareCoord: %d", coord, compareCoord)

		if coord == compareCoord+3 {
			return true
		}
	} else if next == curr-1 {
		// Coming from right
		coord := next % lineLength
		compareCoord := compareMe % lineLength
		if coord == compareCoord-3 {
			return true
		}
	} else if next == curr+lineLength {
		// coming from top
		coord := next / lineLength
		compareCoord := compareMe / lineLength
		if coord == compareCoord+3 {
			return true
		}
	} else if next == curr-lineLength {
		// coming from bottom
		coord := next / lineLength
		compareCoord := compareMe / lineLength
		if coord == compareCoord-3 {
			return true
		}
	}

	return false
}

func main() {
	log.SetPrefix(color.Green + "[ # 17 ] " + color.Reset)
	log.SetFlags(0)

	// The supplied input can be interpreted as a 2D grid of node weights, where
	// we want to find a path with the lowest weight. Apart from the limitation
	// on the number of consecutive steps in a single direction, this is a
	// shortest path problem ,so time to break out Dijkstra's algorithm.
	adjs, lines := readInput("example.txt")

	lineLength := len(lines[0])

	// Fill dists with near-as-makes-no-difference infinity
	dists := make([]int, len(adjs))
	seen := make([]bool, len(adjs))
	prev := make([]int, len(adjs))

	for i := range adjs {
		dists[i] = largeInt
		seen[i] = false
		prev[i] = -1
	}

	source := 0
	target := len(adjs) - 1

	dists[source] = 0

	log.Printf("Source: %d, Target: %d", source, target)

	curr := source

	for curr != -1 {
		curr = getLowestUnvisited(seen, dists)
		if curr == -1 {
			continue
		}

		// log.Printf("Current node: %d", curr)
		seen[curr] = true

		for _, edge := range adjs[curr] {
			if !seen[edge.to] {
				dist := dists[curr] + edge.weight

				// Make sure this is not the 4th step in a row in the same direction
				isFourthStep := isFourthMove(curr, lineLength, edge.to, prev)
				if isFourthStep {
					log.Printf("Skipping node[%d] because it's the 4th step in a row", edge.to)
					continue
				}

				if dist < dists[edge.to] {
					log.Printf("setting node[%d] to %d (was %d)", edge.to, dist, dists[edge.to])
					dists[edge.to] = dist
					prev[edge.to] = curr
				}
			}
		}
	}

	pathNode := target
	path := []int{pathNode}

	for prev[pathNode] != -1 {
		pathNode = prev[pathNode]
		path = append(path, pathNode)
	}

	// Reverse the path
	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	log.Printf("Path: %v", path)
	printGrid(lines, path)

	totalDistance := dists[target]
	log.Printf("Total distance: %d", totalDistance)
}
