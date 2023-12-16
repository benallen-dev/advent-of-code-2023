package main

import (
	"log"
	"sort"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func part01(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		// if i < j return true
		a := hands[i]
		b := hands[j]

		if a.Variant(1) <  b.Variant(1) {
			return true
		}

		if a.Variant(1) > b.Variant(1) {
			return false
		}

		// Variants are equal
		for i := 0; i < 5; i++ {
			cardA := string(a.cards[i])
			cardB := string(b.cards[i])

			valueA := scoreMap[cardA]
			valueB := scoreMap[cardB]

			if valueA != valueB {
				return valueA < valueB
			}
		}

		return false
	})

	total := 0
	for idx, hand := range hands {
		total += (idx + 1) * hand.bid
	}

	log.Println("Total part 1:", total)
}

// This isn't dry but it's a toy program
func part02(hands []Hand) {

	sort.Slice(hands, func(i, j int) bool {
		// if i < j return true
		a := hands[i]
		b := hands[j]

		if a.Variant(2) <  b.Variant(2) {
			return true
		}

		if a.Variant(2) > b.Variant(2) {
			return false
		}

		// Variants are equal
		for i := 0; i < 5; i++ {
			cardA := string(a.cards[i])
			cardB := string(b.cards[i])

			valueA := scoreMapPartTwo[cardA]
			valueB := scoreMapPartTwo[cardB]

			if valueA != valueB {
				return valueA < valueB
			}
		}

		return false
	})

	total := 0
	for idx, hand := range hands {
		total += (idx + 1) * hand.bid
	}

	log.Println("Total part 2:", total)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 07 ] " + color.Reset)

	hands := readInput("input.txt")

	part01(hands)
	part02(hands)
}
