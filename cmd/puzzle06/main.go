package main

import (
	"log"
	"sync"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// It's 4 pairs, so why bother reading from a file
var input1 = [][]int{
	{54, 239},
	{70, 1142},
	{82, 1295},
	{75, 1253},
}

var input2 = [][]int{
	{54708275, 239114212951253},
}

func countWinPermutations(time int, record int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	wins := 0

	for i := 0; i < time; i++ {
		speed := i
		moveTime := time - i

		distance := speed * moveTime

		if distance > record {
			wins++
		}
	}

	resultChan <- wins
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 06 ] " + color.Reset)

	resultChan := make(chan int)
	wg := sync.WaitGroup{}

	// Toggle between input1 and input2 to run the different parts
	input := input2

	for i := 0; i < len(input); i++ {
		time := input[i][0]
		record := input[i][1]

		wg.Add(1)
		go countWinPermutations(time, record, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	margin := 1
	for result := range resultChan {
		margin *= result
	}

	log.Printf("Margin: %d", margin)

}


