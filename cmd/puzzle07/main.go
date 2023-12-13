package main

import (
	"log"
	"sort"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func compareScore(a, b Score) int {
	if a.variant > b.variant {
		return 1
	} else if a.variant < b.variant {
		return -1
	}

	for i := 0; i < len(a.cardOrder); i++ {
		if scoreMap[a.cardOrder[i]] > scoreMap[b.cardOrder[i]] {
			return 1
		} else if scoreMap[a.cardOrder[i]] < scoreMap[b.cardOrder[i]] {
			return -1
		}
	}

	return 0
}

func part01(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		return compareScore(hands[i].Score(), hands[j].Score()) == 1
	})

	total := 0
	for idx, hand := range hands {
		log.Println(hand.Score())

		total += (idx + 1) * hand.bid
	}



	// sort hands based on power
	// for each hand
		// value = (idx +1) * bid
		// total += value
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 07 ] " + color.Reset)

	log.Println("Hello from day 7")

	hands := readInput("input.txt") 

	part01(hands)
}
