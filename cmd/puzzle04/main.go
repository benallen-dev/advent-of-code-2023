package main

import (
	"log"
	"math"
	"sync"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var (
	DEBUG = false
)

func getMatches(card ScratchCard) int {

	matches := 0

	for _, winningNumber := range card.winningNumbers {
		for _, cardNumber := range card.cardNumbers {
			if winningNumber == cardNumber {
				matches += 1
				break // no need to check the rest of the card numbers
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

	// Part two, run through the queue
	// cardCount++ and append winning cards, until queue is empty
	log.Println(q)

	cardCount := 0
	var mu sync.Mutex

	for len(q) > 0 {
		// Do this mutex thing so I'm sure stuff happens in order
		mu.Lock()
		// Increment cardCount
		cardCount++
		// Pop the first card from the queue
		cardId := q[0]
		q = q[1:]

		// Find the card with that id
		card := cards[cardId-1] // off by 1 because card ids start at 1

		matches := getMatches(card)

		if matches > 0 {
			for i := 1; i <= matches; i++ {
				newCardId := cardId + i
				q = append(q, newCardId)
			}
		}

		log.Println("queue length:", len(q))
		mu.Unlock()
	}

	log.Println("Part two:", cardCount)

}
