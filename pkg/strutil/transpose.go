package strutil

import (
	"strings"
)

// Transpose takes a slice of strings and returns a slice of strings where each string is a column of the original
func Transpose(input []string) []string {
	output := strings.Split(input[0], "")

	for i, line := range input {
		if i == 0 {
			continue
		}

		for j, char := range line {
			output[j] += string(char)
		}
	}

	return output
}
