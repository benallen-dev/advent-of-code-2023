package main

import (
//	"log"
	"fmt"
)

type Hashmap map[int]map[string]Lens

func (h Hashmap) String() string {
	str := "\n"

	// Let's build an array of all buckets
	for i := 0; i < 256; i++ {
		bucket, ok := h[i]
		if !ok || len(bucket) == 0 {
			continue
		}

		str += fmt.Sprintf("Box %d: ", i)

		// Lens array
		lenses := make([]Lens, len(bucket))
		for _, lens := range bucket {
			lenses[lens.position] = lens
		}

		for _, lens := range lenses {
			str += fmt.Sprintf("%d:[%s %d] ", lens.position, lens.label, lens.focalLength)
		}
	}


	return str
}
