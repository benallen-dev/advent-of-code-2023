package main

import (
	"fmt"
	"sync"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
	"github.com/benallen-dev/advent-of-code-2023/pkg/strutil"
)

var (
	mtx sync.Mutex
	colors = []string{color.Red, color.Green, color.Yellow, color.Blue, color.Purple, color.Cyan, color.Gray}
)

func printPattern(pattern []string, matchingLines [][]int) {
	mtx.Lock()
	colorMap := map[int]string{}

	// Get colour map
	for i, linePair := range matchingLines {
		colorMap[linePair[0]] = colors[i % len(colors)]
		colorMap[linePair[1]] = colors[i % len(colors)]
	}

	fmt.Println()
	fmt.Println("   1234567890123456789")
	for i, line := range pattern {
		if lineColor, ok := colorMap[i]; ok {
			fmt.Println(fmt.Sprintf("%2d ", i + 1) + lineColor + line + color.Reset)
		} else {
			fmt.Println(fmt.Sprintf("%2d ", i + 1) + line)
		}
	}

	fmt.Println()
	mtx.Unlock()
}

func printPatternTransposed(pattern []string, matchingColumns [][]int) {
	mtx.Lock()
	colorMap := map[int]string{}

	untransposedPattern := strutil.Transpose(pattern)

	// Get colour map
	for i, colPair := range matchingColumns {
		colorMap[colPair[0]] = colors[i % len(colors)]
		colorMap[colPair[1]] = colors[i % len(colors)]
	}

	fmt.Println()
	fmt.Println("   1234567890123456789")
	for i, line := range untransposedPattern {
		fmt.Print(fmt.Sprintf("%2d ", i + 1))
		for j, char := range line {
			if colColor, ok := colorMap[j]; ok {
				fmt.Print(colColor + string(char) + color.Reset)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	mtx.Unlock()
}
