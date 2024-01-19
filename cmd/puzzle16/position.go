package main

import (
	"fmt"
)

type Position [2]int

func (p Position) String() string {
	return fmt.Sprintf("%d", p[0]) + "," + fmt.Sprintf("%d", p[1])
}

func (p Position) Diff(other Position) Position {
	return Position{p[0] - other[0], p[1] - other[1]}
}
