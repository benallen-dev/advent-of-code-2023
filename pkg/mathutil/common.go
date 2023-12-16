package mathutil

// Greatest Common Divider is the largest number that divides both a and b without leaving a remainder
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Lowest Common Multiple is the smallest number that is a multiple of both a and b
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

