package main

import (
	"log"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func getMaxFromPlays(plays []Play) Play {
	max := Play{}

	for _, play := range plays {
		if play.r > max.r {
			max.r = play.r
		}
		if play.g > max.g {
			max.g = play.g
		}
		if play.b > max.b {
			max.b = play.b
		}
	}

	return max
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 02 ] " + color.Reset)

	total := 0
	powerTotal := 0

	games := readInput("input.txt")
	cubes := Play{
		r: 12,
		g: 13,
		b: 14,
	}

	for _, game := range games {
		max := getMaxFromPlays(game.plays)
	
		// Part 1, total of possible within the limits of the 'cubes' play
		if max.r <= cubes.r && max.g <= cubes.g && max.b <= cubes.b {
			total += game.id
		}

		// Part 2, compute power of each game's max cube counts
		power := max.r * max.g * max.b
		powerTotal += power
	}

	log.Printf("Total from possible games (part 1): %d", total)
	log.Printf("Total power from all games (part 2): %d", powerTotal)
}
