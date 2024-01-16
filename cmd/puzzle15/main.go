package main

import (
	"strconv"
	"log"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var (
	bigAssHashmap Hashmap = Hashmap{}
)

// Because these are in a hashmap I guess I gotta do things to em
type Lens struct {
	position int
	power    int
	label    string
}

func main() {
	log.SetPrefix(color.Green + "[ # 15 ] " + color.Reset)
	log.SetFlags(0)

	steps := readInput("example.txt")

	total := 0

	for _, step := range steps {
		total += hash(step)
	}

	log.Println("Part 1:", total)
	log.Println()

	// Ok let's do the hashmap thing
	// We have 256 buckets
	// Each bucket contains another 'map' of strings to ints, but ORDERED, so we need to store positions

	for _, step := range steps {
		removeLens := strings.Contains(step, "-")
		upsertLens := strings.Contains(step, "=")

		label := ""
		power := 0 // special case because they are actually 1-indexed - a lens power of 0 indicates a problem

		if removeLens {
			label = strings.Split(step, "-")[0]
		} else if upsertLens {
			parts := strings.Split(step, "=")

			// Because upsert means we know 100% that there's a = in the string we can ignore the error
			power, _ = strconv.Atoi(parts[1])
			label = parts[0]
		}

		// Now we have the label, we can get the box # by hashing it
		// If the box doesn't exist, create it
		boxNumber := hash(label)

		if removeLens && !upsertLens {
			label = strings.Split(step, "-")[0]

			delete(bigAssHashmap[boxNumber], label)

			// Now we need to reposition the lenses

			log.Fatal("TODO: reposition lenses")
		} else if !removeLens && upsertLens {
			parts := strings.Split(step, "=")

			// Because upsert means we know 100% that there's a = in the string we can ignore the error
			power, _ = strconv.Atoi(parts[1])
			label = parts[0]

			// If the box doesn't exist, create it
			if _, ok := bigAssHashmap[boxNumber]; !ok {
				bigAssHashmap[boxNumber] = map[string]Lens{}
				bigAssHashmap[boxNumber][label] = Lens{label: label, power: power, position: 0}
			}

			// If the box does exist, check if the label exists
			//	If it does, update the power
			//	If it doesn't, add it
			if _, ok := bigAssHashmap[boxNumber][label]; ok {
				// I'm not sure why I have to reference it like this but ok
				lens := bigAssHashmap[boxNumber][label]
				lens.power = power
			} else {
				bigAssHashmap[boxNumber][label] = Lens{label: label, power: power, position: len(bigAssHashmap[boxNumber])}
			}
		}
	}

	// For each box, gather the lensen, get their FocusPower, multiply by their position
	totalPower := 0
	for boxNumber, lenses := range bigAssHashmap {

		for _, lens := range lenses {
			boxValue := boxNumber + 1
			power := boxValue * (lens.position + 1) * lens.power

			log.Printf("%s: %d (box %d) * %d (slot) * $d (focal length) = %d", lens.label, boxValue, lens.position, lens.power, power)
			totalPower += power
		}
	}



}
