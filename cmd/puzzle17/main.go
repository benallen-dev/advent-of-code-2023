package main

import (
	"log"
	"strconv"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

type GraphEdge struct {
	to int
	weight int
}

func createEdge(line int, col int, input []string) GraphEdge {
	// This function fully assumes you're not passing out of bounds lines or columns

	number := (line * len(input[0]) + col)
	weight, err := strconv.Atoi(string(input[line][col]))
	if err != nil {
		log.Panic("Cannot convert weight to int", err)
	}

	return GraphEdge{number, weight}
}

func buildAdjecencyList(input []string) [][]GraphEdge {
	out := make([][]GraphEdge, len(input[0]) * len(input))

	for i, line := range input {
		for j := range line {
			// This is a node
			nodeNumber := (i * len(input[0])) + j

			out[nodeNumber] = []GraphEdge{}

			// It has four potential edges
			if i > 0 {
				newEdge := createEdge(i-1, j, input)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}

			if i < len(input) - 1 {
				newEdge := createEdge(i+1, j, input)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}

			if j > 0 {
				newEdge := createEdge(i, j-1, input)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}

			if j < len(input[0]) - 1 {
				newEdge := createEdge(i, j+1, input)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}
		}
	}	

	return out
}	

func main() {
	log.SetPrefix(color.Green + "[ # 17 ] " + color.Reset)
	log.SetFlags(0)


	input := readInput("example.txt")

	for i, line := range input {
		log.Printf("%2d : %s", i, line)
	}

	// Each character is a node, which we'll label incrementally starting 
	// with the first line, left to right, top to bottom, much like counting
	// words in a book.
	adjs := buildAdjecencyList(input)

	for i, adj := range adjs {
		log.Printf("%2d : %v", i, adj)
	}



}
