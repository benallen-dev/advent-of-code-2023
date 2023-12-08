package main

// mapValue takes a value and a map and returns the mapped value if it exists, otherwise returns the original value
func mapValue(value int, valueMap map[int]int) int{
	mapValue, ok := valueMap[value]
	if !ok {
		return value
	} else {
		return mapValue
	}
}
