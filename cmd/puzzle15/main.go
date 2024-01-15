package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetPrefix(color.Green + "[ # 15 ] " + color.Reset)
	log.SetFlags(0)

	steps := readInput("input.txt")

	for _, step := range steps {
		log.Println(step)
	}
}
