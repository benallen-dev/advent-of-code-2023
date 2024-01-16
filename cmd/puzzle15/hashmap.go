package main

import (
	"fmt"
)

type Hashmap map[int]map[string]Lens

func (h Hashmap) String() string {
	str := ""

	for bucket, lensMap := range h {
		str += fmt.Sprintf("Box %d: ", bucket)

		// Lens array
		lenses := make([]Lens, len(lensMap))
		for _, lens := range lensMap {
			lenses[lens.position] = lens
		}

		for _, lens := range lenses {
			str += fmt.Sprintf("[%s %d] ", lens.label, lens.power)
		}

		str += "\n"
	}


	return str
}
