
package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var instructions = "LRRRLRRRLRRLRLRRLLRRLLRLRRRLRRLRRRLRRLLRLRLRRRLRLLRRRLLRLRRRLRLRRRLRRRLRRRLRRRLRLLLRRRLRRLRRLRRRLRLRLRRLRLRRRLRLRLRLRRRLRRLRLRRRLRRLRRRLRRLLRRRLLRLLRLRRRLRLLRRLLRRRLRLLRRLRLRRLRRRLRLRLRLLRLRRRLRRRLRLLLRRRLRLRRRLRRLRRLLLLRLRRRLRLRRRLLRRRLRRRLRRRLLLRLRLRLLLLRRRLRRLRRRLRLRLRLRRRLRLRRRR"

func part01(nodes map[string]Node) {
	steps := 0
	curentNode := nodes["AAA"]
	
	for curentNode.name != "ZZZ" {
		instruction := instructions[steps % len(instructions)]
		steps++

		if instruction == 'L' {
			curentNode = nodes[curentNode.left]
		} else {
			curentNode = nodes[curentNode.right]
		}
	}

	log.Println("Part 1:", steps)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 07 ] " + color.Reset)

	nodes := readInput("input.txt")

	part01(nodes)
}
