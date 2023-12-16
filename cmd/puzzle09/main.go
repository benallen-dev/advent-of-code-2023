package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func intervals (elements []int) []int {
	result := []int{}

	for i := 1; i < len(elements); i++ {
		result = append(result, elements[i] - elements[i - 1])
	}

	return result
}

func predictNext (elements []int) int {
	// Implement the algorithm explained in the puzzle

	// Part 2 is gonna be brutal if they explain step by step
	// how to solve part 1
	return -1
}

func main() {
	log.SetPrefix(color.Green + "[ # 09 ] " + color.Reset)
	log.SetFlags(0)

	log.Println("Hello from puzzle 09")

	histories := readInput("exampleInput.txt")

	firstExample := histories[0]

	log.Println(firstExample)
	log.Println(intervals(firstExample))
}

