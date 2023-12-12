package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(filename string) []Hand {
	fileBytes, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}
	
	inputStrings := strings.Split(string(fileBytes), "\n") // The last line is blank
	inputStrings = inputStrings[:len(inputStrings)-1]
	hands := []Hand{}

	for _, line := range inputStrings {
		var cards string
		var bid int
		
		n, err := fmt.Sscanf(line, "%s %d", &cards, &bid)
		if (err != nil || n != 2) {
			log.Panic("Cannot parse line", line)
		}

		hands = append(hands, Hand{cards, bid})
	}

	return hands
}
