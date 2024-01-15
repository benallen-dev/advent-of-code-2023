package main

import (
	"log"
	"math"
	"strings"

	"crypto/sha256"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

var (
	cycleCache = map[[32]byte][]string{}
	history = [][32]byte{}
)

func calculateLoad(foo []string) (load int) {

	for i, line := range foo {
		rockLoad := len(foo) - i
		// count number of O's in line
		for _, char := range line {
			if char == 'O' {
				load += rockLoad
			}
		}
	}

	return load
}

func cycle(pattern []string) (cycled []string, stateHash [32]byte) {
	// Check cache	
	key := sha256.Sum256([]byte(strings.Join(pattern, "")))

	if cached, ok := cycleCache[key]; ok {
		return cached, key
	}

	cycled = shiftNorth(pattern)
	cycled = shiftWest(cycled)
	cycled = shiftSouth(cycled)
	cycled = shiftEast(cycled)

	// Cache the result
	cycleCache[key] = cycled

	return cycled, key
}

func mirrorHorizontal(pattern []string) (mirrored []string) {
	for _, line := range pattern {
		mirrored = append(mirrored, strutil.Reverse(line))
	}

	return mirrored
}

func printPattern(pattern []string) {
	for _, line := range pattern {
		log.Println(line)
	}
}

func main() {
	log.SetPrefix(color.Green + "[ # 14 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("input.txt")

	log.Println("Part 1: ", calculateLoad(shiftNorth(input)))

	cycled := input
	stateHash := [32]byte{}

	loops := 1_000_000_000

	bypassHistory := false // it's gross but whatever
	for i := 0; i < loops; i++ {
		if i%100_000 == 0 && i != 0 {
			percentDone := float32(i) / float32(loops) * 100
			log.Println("Progress:", percentDone, "%")
		}

		cycled, stateHash = cycle(cycled)

		// Check if we've seen this state before
		for ii, hash := range history {
			if !bypassHistory && hash == stateHash && ii > 2 {
				cycleLength := len(history) - ii
				iterationsLeft := 1_000_000_000 - i

				// Find how many of these cycles we can fit in iterationsLeft
				cyclesToRun := math.Floor(float64(iterationsLeft) / float64(cycleLength))

				// This means we can subtract cyclesToRun * cycleLength from iterationsLeft
				iterationsLeft -= int(cyclesToRun) * cycleLength

				// Update i
				i = 1000_000_000 - iterationsLeft 
				bypassHistory = true // Otherwise it's gonna detect loops all over again, also yes this is kinda meh
			}
		}

		history = append(history, stateHash)
	}

	log.Println("Part 2: ", calculateLoad(cycled))

}
