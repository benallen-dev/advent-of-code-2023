package strutil

import (
	"strings"
)

// LeftPad pads a string with a given character to a given length
//
// Example:
// leftPad("hello", " ", 10) // "     hello"
func LeftPad(str, pad string, length int) string {
	return strings.Repeat(pad, length-len(str)) + str
}
