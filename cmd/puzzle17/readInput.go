package main

import (
	"fmt"
	"log"
	"strconv"
	"os"
	"strings"
)

type GraphEdge struct {
	from int
	to int
	weight int
}

func (e GraphEdge) String() string {

	return fmt.Sprintf("%d -> %d (%d)", e.from, e.to, e.weight)
}

func createEdge(line int, col int, input []string, from int) GraphEdge {
	// This function fully assumes you're not passing out of bounds lines or columns

	number := (line * len(input[0]) + col)
	weight, err := strconv.Atoi(string(input[line][col]))
	if err != nil {
		log.Panic("Cannot convert weight to int", err)
	}

	return GraphEdge{
		from: from,
		to: number,
		weight: weight,
	}

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
				newEdge := createEdge(i-1, j, input, nodeNumber)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}

			if i < len(input) - 1 {
				newEdge := createEdge(i+1, j, input, nodeNumber)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}

			if j > 0 {
				newEdge := createEdge(i, j-1, input, nodeNumber)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}

			if j < len(input[0]) - 1 {
				newEdge := createEdge(i, j+1, input, nodeNumber)
				out[nodeNumber] = append(out[nodeNumber], newEdge)
			}
		}
	}	

	return out
}	

func readInput(filename string) [][]GraphEdge {
	fileBytes, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileBytes), "\n")

	out := buildAdjecencyList(lines[:len(lines)-1])
	return out
}
