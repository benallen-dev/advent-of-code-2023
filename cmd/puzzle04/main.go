package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var (
	DEBUG = false
)

func main () {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 04 ] " + color.Reset)

	cards := readInput("input.txt")

	for _, card := range cards {
		log.Println(card)
	}
}
