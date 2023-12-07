package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 05 ] " + color.Reset)

	log.Println("Hello from puzzle 05")

	_ = readInput("input.txt")
}
