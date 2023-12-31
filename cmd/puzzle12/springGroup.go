package main

import (
	"fmt"
)

// SpringGroup is a struct that holds the springs and groups of a given line
//
// springs is a string of . # and ? characters
//   . = undamaged spring
//   # = damaged spring
//   ? = unknown spring
// groups is a slice of integers that represent contiguous groups of damaged springs
//
// Example:
//   .springs = ".#....##.??"
//   .groups = [1, 2, 1]
type SpringGroup struct {
	springs string
	groups []int
}

func (sg SpringGroup) String() string {
	return fmt.Sprintf("%s %v", sg.springs, sg.groups)
}

func (sg SpringGroup) Arrangements() int {
	return -1
}
