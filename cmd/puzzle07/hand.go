package main

import (
	"fmt"
)

type Score struct {
	variant   Variant
	cardOrder []string
}

func (s Score) String() string {
	return fmt.Sprintf("%s %v", s.variant, s.cardOrder)
}

type Hand struct {
	cards string
	bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%s %d", h.cards, h.bid)
}

func (h Hand) SortedCards() []string {
	temp := []int{}

	for _, card := range h.cards {
		temp = append(temp, scoreMap[string(card)])
	}

	for i := 0; i < len(temp); i++ {
		for j := i + 1; j < len(temp); j++ {
			if temp[i] < temp[j] {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
	}

	tempString := []string{}
	for _, card := range temp {
		tempString = append(tempString, scoreMapReverse[card])
	}

	return tempString
}

func (h Hand) Score() Score {
	var variant Variant
	var cardOrder []string

	cardmap := map[string]int{}
	for _, card := range h.cards {
		cardmap[string(card)]++
	}

	// Determine the variant and highest cards
	sortedCards := h.SortedCards()

	switch len(cardmap) {
	case 1:
		variant = FiveOfAKind
		cardOrder = []string{string(h.cards[0])}

	case 2:
		// if there is a 3, it's a full house
		// otherwise it's a four of a kind
		for _, count := range cardmap {
			if count == 3 {
				variant = FullHouse
				cardOrder = []string{sortedCards[0], sortedCards[3]}
				break
			} else {
				variant = FourOfAKind
				cardOrder = []string{sortedCards[0], sortedCards[4]}
			}
		}
	case 3:
		// if there is a 3, it's a three of a kind
		// otherwise it's a two pair
		for _, count := range cardmap {
			if count == 3 {
				variant = ThreeOfAKind
				cardOrder = []string{sortedCards[0], sortedCards[3], sortedCards[4]}
			} else {
				variant = TwoPair
				cardOrder = []string{sortedCards[0], sortedCards[2], sortedCards[4]}
			}
		}
	case 4:
		variant = Pair
		cardOrder = []string{sortedCards[0], sortedCards[2], sortedCards[3], sortedCards[4]}
	default:
		variant = HighCard
		cardOrder = sortedCards
	}

	return Score{variant, cardOrder}
}
