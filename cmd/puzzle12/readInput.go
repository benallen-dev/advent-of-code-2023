package main

import (
	"log"
	"os"
	"strings"
	"strconv"
)

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

		for _, group := range groups {
			integer, err := strconv.Atoi(group)
			if (err != nil) {
				log.Panic("Cannot convert group to integer", err)
			}

			integers = append(integers, integer)
		}

		springGroups = append(springGroups, SpringGroup{parts[0], integers})
	}

	return springGroups
}
