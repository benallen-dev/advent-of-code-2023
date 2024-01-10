package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// Left is .
// Right is #
type SpringGroupTreeNode struct {
	springs string
	left    *SpringGroupTreeNode
	right   *SpringGroupTreeNode
}

func (node SpringGroupTreeNode) String() string {
	output := fmt.Sprintf("%s\n", node.springs)
	output += fmt.Sprintf("L: %v\n", node.left)
	output += fmt.Sprintf("R: %v\n", node.right)
	return output
}

// generateTreeNode takes a string of springs and a slice of groups and returns a SpringGroupTreeNode
//
// Springs represents the new spring arrangement
// Groups represents the groups of damaged springs we need to check against
func generateTreeNode(springs string, groups []int) *SpringGroupTreeNode {
	newNode := SpringGroupTreeNode{springs, nil, nil}

	if !strings.Contains(springs, "?") {
		return &newNode
	}

	// Check groups until the first "?"
	partial, _, found := strings.Cut(springs, "?")
	if !found {
		return &newNode
	}

	// Special case where you can just blindly add both because the first char is a ?
	if len(partial) == 0 {
		newLeftSprings := strings.Replace(springs, "?", ".", 1)  // replace first "?" with "."
		newRightSprings := strings.Replace(springs, "?", "#", 1) // replace first "?" with "#"

		newNode.left = generateTreeNode(newLeftSprings, groups)
		newNode.right = generateTreeNode(newRightSprings, groups)

		return &newNode
	}

	partialGroups, err := getGroupNumbers(partial)
	if err != nil {
		log.Println(color.Yellow + "WARN: " + color.Reset + err.Error())
		return nil
	}

	// If we've made too many groups, nope outta there
	if len(partialGroups) > len(groups) {
		// log.Printf(color.Gray + "DEBUG: " + color.Reset + "Too many groups for %s %v", partial, groups)
		return nil
	}

	lastRune := rune(partial[len(partial)-1])

	// Now let's check the groups one by one
	for idx, partialGroup := range partialGroups {
		// If one of the groups is too big, return nil
		if partialGroup > groups[idx] {
			// log.Printf(color.Gray + "DEBUG: " + color.Reset + "Group %d is too big for %s %v", idx, partial, groups)
			return nil
		}

		// If a group is not the last and is too small, return nil
		if (idx < len(partialGroups)-1) && (partialGroup < groups[idx]) {
			return nil
		}

		if idx == len(partialGroups)-1 && partialGroup < groups[idx] && lastRune != '#' {
			return nil
		}
	}

	// Cool, we're still here, so either we're building a new group, or we can add both a dot and a hash

	// If previous is a dot you can always add both
	if lastRune == '.' {
		newLeftSprings := strings.Replace(springs, "?", ".", 1)  // replace first "?" with "."
		newRightSprings := strings.Replace(springs, "?", "#", 1) // replace first "?" with "#"

		newNode.left = generateTreeNode(newLeftSprings, groups)
		newNode.right = generateTreeNode(newRightSprings, groups)
	} else {
		// The previous rune is assumed to be a "#"

		// If the last partial group is less than the target, add a hash
		if partialGroups[len(partialGroups)-1] < groups[len(partialGroups)-1] {
			newRightSprings := strings.Replace(springs, "?", "#", 1) // replace first "?" with "#"
			newNode.right = generateTreeNode(newRightSprings, groups)
		}

		// If the last partial group is equal to the target, add both
		if partialGroups[len(partialGroups)-1] == groups[len(partialGroups)-1] {
			newLeftSprings := strings.Replace(springs, "?", ".", 1) // replace first "?" with "."
			newNode.left = generateTreeNode(newLeftSprings, groups)
		}
	}
	// Part01 code starts here
	// newLeftSprings := strings.Replace(springs, "?", ".", 1)  // replace first "?" with "."
	// newRightSprings := strings.Replace(springs, "?", "#", 1) // replace first "?" with "#"

	// newNode.left = generateTreeNode(newLeftSprings, groups)
	// newNode.right = generateTreeNode(newRightSprings, groups)

	return &newNode
}
