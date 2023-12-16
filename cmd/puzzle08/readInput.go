package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

var inputRegex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

// If you use a hashmap of nodes, left and right can become strings so you don't have to worry about the order of the input
func readInput(filename string) map[string]Node {
	fileBytes, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	inputStrings := strings.Split(string(fileBytes), "\n")
	inputStrings = inputStrings[:len(inputStrings)-1] // The last line is blank


	nodes := map[string]Node{}

	for _, line := range inputStrings {
		matches := inputRegex.FindStringSubmatch(line)
	
		if len(matches) != 4 {
			log.Panic("Cannot parse line", line)
		}

		name := matches[1]
		left := matches[2]
		right := matches[3]

		nodes[name] = Node{name, left, right}
	}

	return nodes
}
