package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	example     string         = "Game 1: 1 red, 5 blue, 10 green; 5 green, 6 blue, 12 red; 4 red, 10 blue, 4 green"
	gameIdRegex *regexp.Regexp = regexp.MustCompile(`Game (\d+):`)
	blueRegex   *regexp.Regexp = regexp.MustCompile(`(\d+) blue`)
	redRegex    *regexp.Regexp = regexp.MustCompile(`(\d+) red`)
	greenRegex  *regexp.Regexp = regexp.MustCompile(`(\d+) green`)
)

type Play struct {
	r int
	g int
	b int
}

type Game struct {
	id    int
	plays []Play
}

func processColor(playText string, colorRegex *regexp.Regexp) int {
	colorText := colorRegex.FindString(playText)
	
	if colorText == "" {
		return 0
	} else {
		color, err := strconv.Atoi(colorRegex.FindStringSubmatch(playText)[1])
		if err != nil {
			log.Panic("Cannot parse color", err)
		}
		return color
	}
}

func processPlays(playText string) []Play {
	plays := []Play{}

	playsText := strings.Split(playText, ";")

	for _, playText := range playsText {
		log.Printf("\nPlay text: %s", playText)

		g := processColor(playText, greenRegex)
		r := processColor(playText, redRegex)
		b := processColor(playText, blueRegex)

		plays = append(plays, Play{
			r: r,
			g: g,
			b: b,
		})

	}

	return plays
}

func readInput(filename string) []Game {

	games := []Game{}

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Panic("Cannot read input file from disk", err)
	}

	lines := strings.Split(string(fileBytes), "\n")

	for _, line := range lines {
		if line == "" { // Last line is empty, and we don't care about parsing empty lines anyway
			continue
		}

		// Parse game ID
		id, err := strconv.Atoi(gameIdRegex.FindStringSubmatch(line)[1])
		if err != nil {
			log.Panic("Cannot parse game ID", err)
		}

		// Extract plays from line
		gamePlayText := strings.Split(line, ":")[1]
		var plays []Play = processPlays(gamePlayText) // I'm being explicit so I don't forget what the type of plays is

		// Add game to games slice
		games = append(games, Game{
			id:    id,
			plays: plays,
		})
	}

	return games
}
