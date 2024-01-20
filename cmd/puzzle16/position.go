package main

import (
	"fmt"
)

// Position is a 2D coordinate [LINE, COLUMN]
type Position [2]int

func (p Position) String() string {
	return fmt.Sprintf("%d", p[0]) + "," + fmt.Sprintf("%d", p[1])
}

// Continue returns a new Position with the same direction as the diff with previous
func (p Position) Continue(other Position) Position {
	
	diff := Position{p[0] - other[0], p[1] - other[1]}
	result := Position{p[0] + diff[0], p[1] + diff[1]}
	return result
}
