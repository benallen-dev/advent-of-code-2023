package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) [][]int {
	fileContents, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileContents), "\n")
	lines = lines[:len(lines)-1] // Remove the last line because it's empty

	var histories [][]int

	for idx, line := range lines {
		numbers := strings.Split(line, " ")
		histories = append(histories, []int{})
		
		for _, number := range numbers {
			number, err := strconv.Atoi(number)
			if (err != nil) {
				log.Panic("Cannot parse line", line)
			}

			histories[idx] = append(histories[idx], number)
		}
	}

	return histories

}
