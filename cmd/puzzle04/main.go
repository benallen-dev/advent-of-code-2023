package main

import (
	"log"
	"math"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func getMatches(card ScratchCard) int {
	matches := 0

	for _, winningNumber := range card.winningNumbers {
		for _, cardNumber := range card.cardNumbers {
			if winningNumber == cardNumber {
				matches += 1
			}
		}
	}

	return matches
}

func getPoints(matches int) int {
	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, math.Max(float64(matches-1), 0)))
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 04 ] " + color.Reset)

	cards := readInput("input.txt")

	q := []int{}
	points := 0

	for _, card := range cards {
		// Part one, for each card identify how many winning numbers appear in the card numbers
		points += getPoints(getMatches(card))

		// Prepare for part two
		q = append(q, card.id)
	}

	log.Println("Part one:", points)
	
	// Part two, run through the queue, incrementing cardCount and appending winners to the queue
	cardCount := 0

	// Go doesn't have a while loop, except it does
	for len(q) > 0 {
		cardCount++

		// Pop the first card from the queue
		cardId := q[0]
		q = q[1:]

		// Find the card with that id
		card := cards[cardId-1] // off by 1 because card ids start at 1

		matches := getMatches(card)

		if matches > 0 {
			// More off-by-one tomfoolery
			for i := 1; i <= matches; i++ {
				newCardId := cardId + i
				q = append(q, newCardId)
			}
		}

	}

	// Turns out that the mutex wasn't necessary, but it also doesn't slow things
	// down too much. The issue actually stems from having print statements inside
	// the while loop, that makes everything really slow!
	log.Println("Part two:", cardCount)
}
