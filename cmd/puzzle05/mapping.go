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
