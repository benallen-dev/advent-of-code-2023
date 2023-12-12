package main 

import (
	"fmt"

	"github.com/benallen-dev/advent-of-code-2023/cmd/puzzle07/handtype"
)


type Hand struct {
	cards string
	bid int
}

func (h Hand) String() string {
	return fmt.Sprintf("%s %d\n", h.cards, h.bid)
}

