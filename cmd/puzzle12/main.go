package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetPrefix(color.Green + "[ # 12 ] " + color.Reset)
	log.SetFlags(0)

	input := readInput("example.txt")

	for _, line := range input {
		log.Println(line)
	}
}
