package main

import (
	"log"
	"os"
	"strings"
	"strconv"
)

// SpringGroup is a struct that holds the springs and groups of a given line
//
// springs is a string of . # and ? characters
//   . = undamaged spring
//   # = damaged spring
//   ? = unknown spring
// groups is a slice of integers that represent contiguous groups of damaged springs
//
// Example:
//   .springs = ".#....##.??"
//   .groups = [1, 2, 1]
type SpringGroup struct {
	springs string
	groups []int
}


func readInput(filename string) []SpringGroup {
	fileContents, err := os.ReadFile(filename)
	if (err != nil) {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileContents), "\n")
	lines = lines[:len(lines)-1] // Remove the last line because it's empty

	springGroups := []SpringGroup{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		groups := strings.Split(parts[1], ",")
		integers := []int{}

		for i, group := range groups {
			integer, err := strconv.Atoi(group)
			if (err != nil) {
				log.Panic("Cannot convert group to integer", err)
			}

			integers[i] = integer
		}

		springGroups = append(springGroups, SpringGroup{parts[0], integers})
	}

	return springGroups
}
