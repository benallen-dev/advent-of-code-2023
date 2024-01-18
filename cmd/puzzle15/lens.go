package main

import (
	"fmt"
)

type Lens struct {
	position int
	focalLength    int
	label    string
}

func (l Lens) String() string {
	return fmt.Sprintf("%s: %d - pos %d", l.label, l.focalLength, l.position)
}
