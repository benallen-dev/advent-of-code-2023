package main

import (
	"fmt"
)

type Hand struct {
	cards string
	bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%s %d", h.cards, h.bid)
}

func (h Hand) Variant(day int) Variant {
	var variant Variant

	cardmap := map[string]int{}
	for _, card := range h.cards {
		cardmap[string(card)]++
	}

	if day == 2 {
		jokers := cardmap["J"]
		if jokers > 0 {
			// remove the jokers from the map
			delete(cardmap, "J")

			// Find the highest card count
			highestCard := "J" // Js are the lowest card now
			highestCount := 0
			for k, v := range cardmap {
				// if the count is higher, replace the highest card
				if v > highestCount {
					highestCard = k
					highestCount = v
				}
				
				//if the count is equal, only replace the highest card if the value is higher
				if v == highestCount {
					if scoreMapPartTwo[k] > scoreMapPartTwo[highestCard] {
						highestCard = k
						highestCount = v
					}
				}
			}

			// We now know the highest card and its value
			// Set the count of the highest card to itself + the jokers
			cardmap[highestCard] += jokers
		}
	}

	switch len(cardmap) {
	case 1:
		variant = FiveOfAKind
	case 2:
		// if there is a 3, it's a full house
		// otherwise it's a four of a kind
		for _, count := range cardmap {
			if count == 3 {
				variant = FullHouse
				break
			} else {
				variant = FourOfAKind
			}
		}
	case 3:
		// if there is a 3, it's a three of a kind
		// otherwise it's a two pair
		for _, count := range cardmap {
			if count == 3 {
				variant = ThreeOfAKind
				break
			} else {
				variant = TwoPair
			}
		}
	case 4:
		variant = Pair
	default:
		variant = HighCard
	}

	return variant
}
