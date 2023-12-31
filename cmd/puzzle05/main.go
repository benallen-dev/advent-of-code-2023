package main

import (
	"log"
	"sync"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

func mapSeedToLocation(seed int, maps Maps) int {
	soil := maps.seedToSoilMap.getMapValue(seed)
	fertilizer := maps.soilToFertilizerMap.getMapValue(soil)
	water := maps.fertilizerToWaterMap.getMapValue(fertilizer)
	light := maps.waterToLightMap.getMapValue(water)
	temperature := maps.lightToTemperatureMap.getMapValue(light)
	humidity := maps.temperatureToHumidityMap.getMapValue(temperature)
	location := maps.humidityToLocationMap.getMapValue(humidity)

	return location
}

func part01(seeds []int, maps Maps) int {

	lowestLocation := 1000000000000000000 // just some big number

	for _, seed := range seeds {
		location := mapSeedToLocation(seed, maps)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func processRange(start int, end int, maps Maps, resultChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	lowestLocation := 1000000000000000000 // just some big number again
		for j := start; j <= end; j++ {
			location := mapSeedToLocation(j, maps)
			if location < lowestLocation {
				lowestLocation = location
			}
		}

	resultChannel <- lowestLocation
}

func part02(seeds []int, maps Maps) int {

	// Now the seeds are actually tuples of [start, range]
	// So the seed array becomes exponentially larger
	// The way to solve it is to iterate and just discard results as you go
	// instead of storing them all in memory up front

	resultChannel := make(chan int)
	var wg sync.WaitGroup

	lowestLocation := 1000000000000000000 // just some big number again

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := seeds[i] + seeds[i+1]

		log.Printf("Processing seed range %d to %d", start, end)

		wg.Add(1)
		go processRange(start, end, maps, resultChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		log.Printf("Got result: %d", result)
		if result < lowestLocation {
			lowestLocation = result
		}
	}

	return lowestLocation
}

func main() {

	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 05 ] " + color.Reset)

	puzzleData := readInput("input.txt")
	// puzzleData := readInput("exampleInput.txt")

	seeds := puzzleData.seeds
	maps := buildMaps(puzzleData)

	// Part 01
	lowestLocation := part01(seeds, maps)
	log.Printf("Lowest location part 01: %d", lowestLocation)

	// Part 02
	lowestLocation = part02(seeds, maps)
	log.Printf("Lowest location part 02: %d", lowestLocation)
}
