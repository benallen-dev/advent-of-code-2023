package main

type Variant int
const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h Variant) String() string {
	return [...]string{"High Card", "Pair", "Two Pair", "Three of a Kind", "Full House", "Four of a Kind", "Five of a Kind"}[h]
}
