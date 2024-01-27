package main

import (
	"fmt"
)

type GraphEdge struct {
	from int
	to int
	weight int
}

func (e GraphEdge) String() string {

	return fmt.Sprintf("%d -> %d (%d)", e.from, e.to, e.weight)
}

