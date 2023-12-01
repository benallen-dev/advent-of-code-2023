package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetPrefix(color.Green + "[ # 01 ] " + color.Reset)
	log.SetFlags(0)

	testInput := "1fdatwofda3fds4fsdfivefds6fdseveneight9zero"

	// It occurs to me that if I split all these operations into functions
	// I could parralelise all this by starting a goroutine for each line
	// thereby improving performance by a factor however many threads we
	// get to use

	foo := parseLine(testInput)
	foo = collectNumbers(foo)

	log.Println(foo)

	// For each line of the input
	//	1. Replace "one", "two" with "1", "2" etc
	//	2. Collect all the numbers
	//	3. Take the first and last ones and smush em together
	//	4. Convert this string to Int
	//	5. Sum all the numbers

}
