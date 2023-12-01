package main

import (
	"log"
	"regexp"
	"strconv"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var regexNotNumber = regexp.MustCompile(`[^0-9]`)

func main() {
	log.SetPrefix(color.Green + "[ # 01 ] " + color.Reset)
	log.SetFlags(0)

	// It occurs to me that if I split all these operations into functions
	// I could parralelise all this by starting a goroutine for each line
	// thereby improving performance by a factor however many threads we
	// get to use, but I also need to synchronously add to sum, which means
	// mutexes (I think) and I don't want to get into those right now

	input := readInput("input.txt")
	sum := int64(0)

	for i, line := range input {
		if (line == "") { // must be the last one
			continue
		}

		// Replace written numbers with digits
		// TODO: sliding window
		
		// Filter out all results that aren't digits
		foo := regexNotNumber.ReplaceAllLiteralString(line, "")
		if (len(foo) == 0) {
			log.Printf(color.Red + "[WARN] " + color.Reset + "Line #%i: '%s' does not contain numbers", i, foo)
			continue
		}



		// Concat the first and last digits together
		concatted := string(foo[0]) + string(foo[len(foo)-1])

		// Cast from string to int
		fooInt, err := strconv.ParseInt(concatted, 10, 64)
		if err != nil {
			log.Panicf("Could not convert %s to string", foo)
		}

		// Add to running total
		sum += fooInt
	}

	log.Println(sum)
}
