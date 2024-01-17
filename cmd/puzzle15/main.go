package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var (
	bigAssHashmap Hashmap
)

func init() {
	bigAssHashmap = Hashmap{}

	for i := 0; i < 256; i++ {
		bigAssHashmap[i] = map[string]Lens{}
	}
}

func main() {
	log.SetPrefix(color.Green + "[ # 15 ] " + color.Reset)
	log.SetFlags(0)

	steps := readInput("input.txt")

	total := 0

	for _, step := range steps {
		total += hash(step)
	}

	log.Println("Part 1:", total)

	// Ok let's do the hashmap thing
	// We have 256 buckets
	// Each bucket contains another 'map' of strings to ints, but ORDERED, so we need to store positions

	for _, step := range steps {
		removeLens := strings.Contains(step, "-")
		upsertLens := strings.Contains(step, "=")

		// Get the label and focalLength
		label := ""
		focalLength := 0 // special case because they are actually 1-indexed - a lens focalLength of 0 indicates a problem

		if removeLens {
			label = strings.Split(step, "-")[0]
		} else if upsertLens {
			parts := strings.Split(step, "=")

			// Because upsert means we know 100% that there's a = in the string we can ignore the error
			focalLength, _ = strconv.Atoi(parts[1])
			label = parts[0]
		}

		// Now we have the label, we can get the box # by hashing it
		boxNumber := hash(label)

		if removeLens {
			// Go to the relevant box
			box := bigAssHashmap[boxNumber]
			// Remove the lens with the given label if it is present
			if _, ok := box[label]; ok {
				// Keep track of where the lens was
				oldLens := box[label]

				delete(box, label)

				// Then, move any remaining lenses forward
				for _, lens := range box {
					if lens.position > oldLens.position {
						lens.position = lens.position - 1

						// It would be nicer to get a pointer and update the value in place
						box[lens.label] = lens
					}
				}
			}
		}

		if upsertLens {
			// Go to the relevant box
			box := bigAssHashmap[boxNumber]

			// If the lens is present, update the focalLength
			if _, ok := box[label]; ok {
				lens := box[label]
				lens.focalLength = focalLength
				// Is this really necessary?
				box[label] = lens
			} else {
				// If the lens is not present, add it to the end of the box
				box[label] = Lens{label: label, focalLength: focalLength, position: len(box)}
			}
		}
	}

	totalPower := 0
	for boxNumber, lenses := range bigAssHashmap {

		for _, lens := range lenses {
			boxValue := boxNumber + 1
			positionValue := lens.position + 1
			focalLength := boxValue * positionValue * lens.focalLength

			totalPower += focalLength
		}
	}

	log.Println("Part 2:", totalPower)
}
