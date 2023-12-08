package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PuzzleData struct {
	seeds                       []int
	seedToSoilString            []string
	soilToFertilizerString      []string
	fertilizerToWaterString     []string
	waterToLightString          []string
	lightToTemperatureString    []string
	temperatureToHumidityString []string
	humidityToLocationString    []string
}

var (
	seedsRegex                 = regexp.MustCompile(`seeds: (.*)\n`)
	seedToSoilRegex            = regexp.MustCompile(`seed-to-soil map:\n((?:\d+ \d+ \d+\n)+)`)
	soilToFerilizerRegex       = regexp.MustCompile(`soil-to-fertilizer map:\n((?:\d+ \d+ \d+\n)+)`)
	fertilizerToWaterRegex     = regexp.MustCompile(`fertilizer-to-water map:\n((?:\d+ \d+ \d+\n)+)`)
	waterToLightRegex          = regexp.MustCompile(`water-to-light map:\n((?:\d+ \d+ \d+\n)+)`)
	lightToTemperatureRegex    = regexp.MustCompile(`light-to-temperature map:\n((?:\d+ \d+ \d+\n)+)`)
	temperatureToHumidityRegex = regexp.MustCompile(`temperature-to-humidity map:\n((?:\d+ \d+ \d+\n)+)`)
	humidityToLocationRegex    = regexp.MustCompile(`humidity-to-location map:\n((?:\d+ \d+ \d+\n)+)`)
)

func getSeeds(input string) []int {
	seedString := seedsRegex.FindStringSubmatch(input)
	seeds := strings.Split(seedString[1], " ")

	var parsedSeeds []int
	for _, seed := range seeds {
		parsedSeed, err := strconv.Atoi(seed)
		if err != nil {
			log.Panic("Cannot convert seed to int", err)
		}
		parsedSeeds = append(parsedSeeds, parsedSeed)
	}

	return parsedSeeds
}


// Takes the input file string and a regex, and returns all mappings for that regex in an array
func getMapStrings(input string, regex *regexp.Regexp) []string {
	numbers := regex.FindStringSubmatch(input)[1]
	numberLines := strings.Split(numbers, "\n")

	return numberLines[:len(numberLines)-1] // truncate last element, which is empty
}

// Reads the input file and returns the puzzle's input data
func readInput(filename string) PuzzleData {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Panic("Cannot read input file from disk", err)
	}

	input := string(fileContents)

	seeds := getSeeds(input)
	seedToSoilStrings := getMapStrings(input, seedToSoilRegex)
	soilToFertilizerStrings := getMapStrings(input, soilToFerilizerRegex)
	fertilizerToWaterStrings := getMapStrings(input, fertilizerToWaterRegex)
	waterToLightStrings := getMapStrings(input, waterToLightRegex)
	lightToTemperatureStrings := getMapStrings(input, lightToTemperatureRegex)
	temperatureToHumidityStrings := getMapStrings(input, temperatureToHumidityRegex)
	humidityToLocationStrings := getMapStrings(input, humidityToLocationRegex)

	return PuzzleData{
		seeds:      seeds,
		seedToSoilString:            seedToSoilStrings,
		soilToFertilizerString:      soilToFertilizerStrings,
		fertilizerToWaterString:     fertilizerToWaterStrings,
		waterToLightString:          waterToLightStrings,
		lightToTemperatureString:    lightToTemperatureStrings,
		temperatureToHumidityString: temperatureToHumidityStrings,
		humidityToLocationString:    humidityToLocationStrings,
	}
}
