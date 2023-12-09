package main

import (
	"fmt"
)

type Domain struct {
	sourceStart int
	targetStart int
	size        int
}

type Map struct {
	domains []Domain
}

type Maps struct {
	seedToSoilMap Map
	soilToFertilizerMap Map
	fertilizerToWaterMap Map
	waterToLightMap Map
	lightToTemperatureMap Map
	temperatureToHumidityMap Map
	humidityToLocationMap Map
}

func (m Map) getMapValue(value int) int {
	for _, domain := range m.domains {
		if value >= domain.sourceStart && value < domain.sourceStart+domain.size {
			return domain.targetStart + (value - domain.sourceStart)
		}
	}
	return value
}

func generateMap(mapping []string) Map {
	var domains []Domain

	for _, mapping := range mapping {
		var sourceStart, targetStart, size int
		fmt.Sscanf(mapping, "%d %d %d", &targetStart, &sourceStart, &size) // Copilot suggested Sscanf, which I didn't know about
		domains = append(domains, Domain{sourceStart, targetStart, size})
	}

	return Map{domains}
}

func buildMaps(puzzleData PuzzleData) Maps {
	seedToSoilMap := generateMap(puzzleData.seedToSoilString)
	soilToFertilizerMap := generateMap(puzzleData.soilToFertilizerString)
	fertilizerToWaterMap := generateMap(puzzleData.fertilizerToWaterString)
	waterToLightMap := generateMap(puzzleData.waterToLightString)
	lightToTemperatureMap := generateMap(puzzleData.lightToTemperatureString)
	temperatureToHumidityMap := generateMap(puzzleData.temperatureToHumidityString)
	humidityToLocationMap := generateMap(puzzleData.humidityToLocationString)

	return Maps{
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	}
}
