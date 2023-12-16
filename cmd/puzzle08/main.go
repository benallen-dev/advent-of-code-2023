package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/mathutil"
)

var instructions = "LRRRLRRRLRRLRLRRLLRRLLRLRRRLRRLRRRLRRLLRLRLRRRLRLLRRRLLRLRRRLRLRRRLRRRLRRRLRRRLRLLLRRRLRRLRRLRRRLRLRLRRLRLRRRLRLRLRLRRRLRRLRLRRRLRRLRRRLRRLLRRRLLRLLRLRRRLRLLRRLLRRRLRLLRRLRLRRLRRRLRLRLRLLRLRRRLRRRLRLLLRRRLRLRRRLRRLRRLLLLRLRRRLRLRRRLLRRRLRRRLRRRLLLRLRLRLLLLRRRLRRLRRRLRLRLRLRRRLRLRRRR"

type zNodePeriod struct {
	zNode  Node
	period int
}

func part01(nodes map[string]Node) {
	steps := 0
	curentNode := nodes["AAA"]

	for curentNode.name != "ZZZ" {
		instruction := instructions[steps%len(instructions)]
		steps++

		if instruction == 'L' {
			curentNode = nodes[curentNode.left]
		} else {
			curentNode = nodes[curentNode.right]
		}
	}

	log.Println("Part 1:", steps)
}

// Ghosts appear to be on a cycle
func findZNodeInterval(nodes map[string]Node, start string) int {
	steps := 0
	curentNode := nodes[start]
	zNodes := []zNodePeriod{}

	interval := 0

	for steps < 10000000000000000 { // Some crazy high number because we're gonna break anyway
		instruction := instructions[steps%len(instructions)]
		steps++

		if curentNode.name[len(curentNode.name)-1] == 'Z' {
			zNodes = append(zNodes, zNodePeriod{curentNode, steps})

			if len(zNodes) >= 3 {
				lastInterval := zNodes[len(zNodes)-2].period - zNodes[len(zNodes)-3].period
				interval = zNodes[len(zNodes)-1].period - zNodes[len(zNodes)-2].period

				if interval == lastInterval {
					break
				}
			}
		}

		if instruction == 'L' {
			curentNode = nodes[curentNode.left]
		} else {
			curentNode = nodes[curentNode.right]
		}
	
	}

	log.Printf("%s -> %s :: %d steps", start, curentNode.name, interval)

	return zNodes[len(zNodes)-1].period - zNodes[len(zNodes)-2].period
}

func part02(nodes map[string]Node) {

	intervals := []int{
		findZNodeInterval(nodes, "PBA"),
		findZNodeInterval(nodes, "QVA"),
		findZNodeInterval(nodes, "VKA"),
		findZNodeInterval(nodes, "AAA"),
		findZNodeInterval(nodes, "LSA"),
		findZNodeInterval(nodes, "VSA"),
	}

	lowestCommonMultiple := 1
	for _, number := range intervals {
		lowestCommonMultiple = mathutil.LCM(lowestCommonMultiple, number)
	}

	log.Println("Part 2:", lowestCommonMultiple)
}


func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 07 ] " + color.Reset)

	nodes := readInput("input.txt")

	part01(nodes)
	part02(nodes)
}
