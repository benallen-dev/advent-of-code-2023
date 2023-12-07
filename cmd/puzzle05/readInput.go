package main

import (
	"log"
	"strconv"
	// "os"
	// "strings"

	"regexp"
)

var mapping map[int64]int64

func readInput(filename string) []string {
	// seedsInput := "79 14 55 13"

	// Example input from problem statement
	seed2soilInput := []string{
		"50 98 2",
		"52 50 48",
	}

	// The issue with this is that you can only read the map one way - I think you need to read it both ways
	seed2soilMapping := map[int]int{}

	for _, mapping := range seed2soilInput {
		mappingRegex := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
		mappingMatch := mappingRegex.FindStringSubmatch(mapping)

		source, err := strconv.Atoi(mappingMatch[1])
		destination, err := strconv.Atoi(mappingMatch[2])
		amount, err := strconv.Atoi(mappingMatch[3])
		if err != nil {
			log.Panic("Cannot convert string to int", err)
		}

		log.Println(source, destination, amount)

		for i := 0; i < amount; i++ {
			seed2soilMapping[source + i] = destination + i
		}
	}

	log.Println(seed2soilMapping)

	// 	fileContents, err := os.ReadFile(filename)
	// 	if err != nil {
	// 		log.Panic("Cannot read input file from disk", err)
	// 	}

	// 	lines := strings.Split(string(fileContents), "\n")
	// 	lines = lines[:len(lines)-1] // Remove the last line because it's empty

	// return lines
	return []string{}
}
