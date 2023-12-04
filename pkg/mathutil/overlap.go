package mathutil

// RangeOverlap returns true if the ranges overlap
func RangeOverlap(start1, end1, start2, end2 int) bool {
	return start1 <= end2 && end1 >= start2
}
