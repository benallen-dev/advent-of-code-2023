package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

var (
	cardRegex           = regexp.MustCompile(`^Card +([0-9]+):`)
	winningNumbersRegex = regexp.MustCompile(`^Card +[0-9]+: +([0-9 ]+ )|`)
	cardNumbersRegex    = regexp.MustCompile(`\|((?: +[0-9]+)+)`)
)

func readInput(filename string) []ScratchCard {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileContents), "\n")
	lines = lines[:len(lines)-1] // Remove the last line because it's empty

	// Let's read this into a list of Cards
	cards := []ScratchCard{}

	for _, line := range lines {
		cardNumber, err := strconv.Atoi(cardRegex.FindStringSubmatch(line)[1])
		if err != nil {
			log.Panic("Cannot convert card number to int: ", err)
		}

		// Feels kinda gross but whatever
		winningNumberStrings := strings.Split(winningNumbersRegex.FindStringSubmatch(line)[1], " ")
		winningNumbers := []int{}
		for _, winningNumberString := range winningNumberStrings {
			if winningNumberString == "" {
				continue
			}

			winningNumber, err := strconv.Atoi(winningNumberString)
			if err != nil {
				log.Panic("Cannot convert winning number to int: ", err)
			}

			winningNumbers = append(winningNumbers, winningNumber)
		}

		cardNumberStrings := strings.Split(cardNumbersRegex.FindStringSubmatch(line)[1], " ")
		cardNumbers := []int{}
		for _, cardNumberString := range cardNumberStrings {
			if cardNumberString == "" {
				continue
			}

			cardNumber, err := strconv.Atoi(cardNumberString)
			if err != nil {
				log.Panic("Cannot convert card number to int: ", err)
			}

			cardNumbers = append(cardNumbers, cardNumber)
		}
		
		cards = append(cards, ScratchCard{
			id:             cardNumber,
			winningNumbers: winningNumbers,
			cardNumbers:    cardNumbers,
		})
	}

	return cards
}
