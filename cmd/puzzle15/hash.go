package main

import (
	// "log"
	// "fmt"
)

func getASCII(input rune) int {
	return int(input)
}

func hash(key string) int {
	value := 0

	for _, char := range key {
		value += getASCII(char)
		value *= 17
		value %= 256
	}

	return value
}

