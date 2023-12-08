package main

import (
	"log"
	"sync"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

var mut sync.Mutex

// Part 01
// Read seeds
// Map seed -> soil
// Map soil -> fertilizer
// Map fertilizer -> water
// Map water -> light
// Map light -> temperature
// Map temperature -> humidity
// Map humidity -> location

// we now have a list of locations
// Note that at this point we want a map from int(see) -> int(location) so we can sort the map by value and get its key
// Perhaps an object

func main() {

	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 05 ] " + color.Reset)

	puzzleData := readInput("input.txt")
	// puzzleData := readInput("exampleInput.txt")

	seedToSoilMap := generateMap(puzzleData.seedToSoilString)
	soilToFertilizerMap := generateMap(puzzleData.soilToFertilizerString)
	fertilizerToWaterMap := generateMap(puzzleData.fertilizerToWaterString)
	waterToLightMap := generateMap(puzzleData.waterToLightString)
	lightToTemperatureMap := generateMap(puzzleData.lightToTemperatureString)
	temperatureToHumidityMap := generateMap(puzzleData.temperatureToHumidityString)
	humidityToLocationMap := generateMap(puzzleData.humidityToLocationString)

	// Part 01
	seedToLocationMap := map[int]int{}

	lowestLocation := 1000000000000000000 // just some big number

	for _, seed := range puzzleData.seeds {
		soil := seedToSoilMap.getMapValue(seed)
		fertilizer := soilToFertilizerMap.getMapValue(soil)
		water := fertilizerToWaterMap.getMapValue(fertilizer)
		light := waterToLightMap.getMapValue(water)
		temperature := lightToTemperatureMap.getMapValue(light)
		humidity := temperatureToHumidityMap.getMapValue(temperature)
		location := humidityToLocationMap.getMapValue(humidity)

		seedToLocationMap[seed] = location

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	log.Printf("Lowest location: %d", lowestLocation)

}
