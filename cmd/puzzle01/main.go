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

	input := readInput("input.txt")

	sum := int64(0)

	for i, line := range input {
		if line == "" {
			continue
		}

		// Parse written digits into digits
		foo := parseWrittenNumbers(line)

		// Filter out all (remaining) results that aren't digits
		foo = regexNotNumber.ReplaceAllLiteralString(foo, "")
		if len(foo) == 0 {
			log.Printf(color.Red+"[WARN] "+color.Reset+"Line #%i: '%s' does not contain numbers", i, foo)
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
