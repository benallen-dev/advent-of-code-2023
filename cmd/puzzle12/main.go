package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetPrefix(color.Green + "[ # 12 ] " + color.Reset)
	log.SetFlags(0)

	springGroups := readInput("example.txt")

	for _, springGroup := range springGroups {
		log.Println(springGroup)
		log.Println(springGroup.Arrangements(), "arrangements")
		log.Println()
	}
}
