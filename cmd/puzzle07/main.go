package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 07 ] " + color.Reset)

	log.Println("Hello from day 7")

	hands := readInput("input.txt") 
	log.Println(hands)
}
