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

type seedChart struct {
	seed         int
	soil         int
	fertilizer   int
	water        int
	light        int
	temperature  int
	humidity     int
	location     int
}

// func populateSeedChart(seed int) seedChart {
// 	soil := mapSeedToSoil(seed)
// 	fertilizer := mapSoilToFertilizer(soil)
// 	water := mapFertilizerToWater(fertilizer)
// 	light := mapWaterToLight(water)
// 	temperature := mapLightToTemperature(light)
// 	humidity := mapTemperatureToHumidity(temperature)
// 	location := mapHumidityToLocation(humidity)

// 	return seedChart{
// 		seed:         seed,
// 		soil:         soil,
// 		fertilizer:   fertilizer,
// 		water:        water,
// 		light:        light,
// 		temperature:  temperature,
// 		humidity:     humidity,
// 		location:     location,
// 	}
// }

func logMapping (name string, mapping []string) {
	mut.Lock()
	log.Println()
	log.Println(name)
	for _, mapping := range mapping {
		log.Println(mapping)
	}
	mut.Unlock()
}

func main() {

	log.SetFlags(0)
	log.SetPrefix(color.Green + "[ # 05 ] " + color.Reset)

	log.Println("Hello from puzzle 05")

	puzzleData := readInput("input.txt")

	log.Println("Seeds:", puzzleData.seeds)
	logMapping("Seed to soil", puzzleData.seedToSoilString)
	logMapping("Soil to fertilizer", puzzleData.soilToFertilizerString)
	logMapping("Fertilizer to water", puzzleData.fertilizerToWaterString)
	logMapping("Water to light", puzzleData.waterToLightString)
	logMapping("Light to temperature", puzzleData.lightToTemperatureString)
	logMapping("Temperature to humidity", puzzleData.temperatureToHumidityString)
	logMapping("Humidity to location", puzzleData.humidityToLocationString)



	log.Println("End of program")
}
