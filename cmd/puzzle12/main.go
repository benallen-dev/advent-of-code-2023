package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// Recursive solution by HyperNeutrino that finally got me unstuck
// This approach is 3 orders of magnitude faster than my naive "just construct them all" approach
// Check him out at https://www.youtube.com/watch?v=g3Ms5e7Jdqo
func count(cfg string, nums []int, debug bool) int {
	if cfg == "" { // If there are no springs left
		if len(nums) == 0 { // No groups left either, so we're done
			return 1
		} else { // There are still groups left to fill so there are no valid arrangements
			return 0
		}
	}

	if len(nums) == 0 { // If there are no groups left
		if strings.ContainsRune(cfg, '#') { // There are still damaged springs left so there are no valid arrangements
			return 0
		} else {
			return 1
		}
	}

	result := 0
	firstRune := rune(cfg[0])

	if strings.ContainsRune(".?", firstRune) { // If the first spring is undamaged or unknown, we can just skip it
		result += count(cfg[1:], nums, debug)
	}

	if strings.ContainsRune("#?", firstRune) { // Making sure we don't go out of bounds was a pain for this part
		groupFitsInConfig := nums[0] <= len(cfg)
		if groupFitsInConfig {
			groupEndIdx := min(nums[0], len(cfg)-1)
			groupContainsNoDots := !strings.ContainsRune(cfg[:nums[0]], '.')
			noHashAfterGroup := nums[0] == len(cfg) || cfg[groupEndIdx] != '#'

			if debug {
				log.Printf("%s %v", cfg, nums)

				log.Printf("groupEndIdx: %d", groupEndIdx)
				log.Printf("groupFitsInConfig: %t", groupFitsInConfig)
				log.Printf("groupContainsNoDots: %t", groupContainsNoDots)
				log.Printf("noHashAfterGroup: %t", noHashAfterGroup)
			}

			if groupFitsInConfig && groupContainsNoDots && noHashAfterGroup {
				startIdx := min(nums[0]+1, len(cfg))
				result += count(cfg[startIdx:], nums[1:], debug)
			}
		}
	}

	if debug {
		log.Printf("%s %v: %d", cfg, nums, result)
	}

	return result
}

func startWorker(workChannel chan SpringGroup, resultChannel chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()

	log.Println("Starting worker", workerId)

	resultCount := 0

	for {
		select {
		case springGroup, ok := <- workChannel:
			if !ok {
				log.Println("Not ok?")
				return
			}

			result := count(springGroup.springs, springGroup.groups, false)
			resultCount++

			log.Printf("[ Worker %d :: %03d ] %d", workerId, resultCount, result)

			resultChannel <- result
		default:
			log.Println("Default case, guess the queue is empty. Killing worker", workerId)
			return
		}
	}
}

func main() {
	log.SetPrefix(color.Green + "[ # 12 ] " + color.Reset)
	log.SetFlags(0)

	springGroups := readInput("input.txt")

	starttime := time.Now()

	totalArrangements := 0
	for _, springGroup := range springGroups {
		totalArrangements += springGroup.Arrangements()
	}

	log.Printf("Total arrangements part 1: %d", totalArrangements)
	log.Printf("Time taken: %s", time.Since(starttime))

	// Now for the painful bit where my ineficcieny is punished
	// LOL doing this filled up 32GB of RAM in like 20 seconds

	// This is just straight-up not going to work because the numbers are too big.
	// The first line alone contains 49 "?" characters, which means 2^49 possibilities.
	// Even if you stop generating the tree once the groups no longer match that's still
	// an astronomical number.

	hyperNeutrinoTime := time.Now()
	totalHyperNeutrinoArrangements := 0
	for _, springGroup := range springGroups {
		totalHyperNeutrinoArrangements += count(springGroup.springs, springGroup.groups, false)
	}
	log.Println()
	log.Printf("Total arrangements part 1 (HyperNeutrino): %d", totalHyperNeutrinoArrangements)
	log.Printf("Time taken: %s", time.Since(hyperNeutrinoTime))

	// Okidoki let's attempt part 2
	unfoldedSpringGroups := unfoldRecords(springGroups)
	totalArrangementsPart2 := 0

	// Let's try some threading and stuff
	// This will undoubtedly cause headaches when I try to add caching to the
	// count function but just for the lulz for now
	resultChannel := make(chan int)
	workChannel := make(chan SpringGroup, len(unfoldedSpringGroups))
	var wg sync.WaitGroup

	// Fill the work channel
	for _, springGroup := range unfoldedSpringGroups {
		workChannel <- springGroup
	}

	// Start the workers
	var NUM_WORKERS = 4

	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(1)
		go startWorker(workChannel, resultChannel, &wg, i)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		totalArrangementsPart2 += result
		log.Printf("[ main            ] %d items remaining", len(workChannel))
	}

	log.Printf("Total arrangements part 2: %d", totalArrangementsPart2)
}
