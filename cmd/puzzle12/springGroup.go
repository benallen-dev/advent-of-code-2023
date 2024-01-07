package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

// Util functions not a member of any struct

// getGroupNumbers takes a SpringGroup.strings and returns a slice of integers that represent contiguous groups of damaged springs
// It will not be happy if you pass it a string containing ? characters
func getGroupNumbers(springs string) ([]int, error) {
	groups := []int{}
	currentGroup := 0
	
	var previousSpring rune

	for idx, spring := range springs {
		if spring == '#' {
			currentGroup++
		} else if spring == '.' {
			if previousSpring == '#' {
				groups = append(groups, currentGroup)
				currentGroup = 0
			}
		} else {
			return nil, fmt.Errorf("getGroupNumbers: invalid spring character: %c", spring)
		}

		if idx == len(springs) - 1 && currentGroup > 0 {
			groups = append(groups, currentGroup)
		}
		
		previousSpring = spring
	}

	return groups, nil
}

// SpringGroup is a struct that holds the springs and groups of a given line
//
// springs is a string of . # and ? characters
//   . = undamaged spring
//   # = damaged spring
//   ? = unknown spring
// groups is a slice of integers that represent contiguous groups of damaged springs
//
// Example:
//   .springs = ".#....##.??"
//   .groups = [1, 2, 1]
type SpringGroup struct {
	springs string
	groups []int
}

func (sg SpringGroup) String() string {
	return fmt.Sprintf("%s %v", sg.springs, sg.groups)
}

// Left is .
// Right is #
type SpringGroupTreeNode struct {
	springs string
	left *SpringGroupTreeNode
	right *SpringGroupTreeNode
}

func (node SpringGroupTreeNode) String() string {
	output := fmt.Sprintf("%s\n", node.springs)
	output += fmt.Sprintf("L: %v\n", node.left)
	output += fmt.Sprintf("R: %v\n", node.right)
	return output
}

func generateTreeNode(springs string) *SpringGroupTreeNode {
	newNode := SpringGroupTreeNode{springs, nil, nil}

	if !strings.Contains(springs, "?") {
		return &newNode
	}

	newLeftSprings := strings.Replace(springs, "?", ".", 1) // replace first "?" with "."
	newRightSprings := strings.Replace(springs, "?", "#", 1) // replace first "?" with "#"

	newNode.left = generateTreeNode(newLeftSprings)
	newNode.right = generateTreeNode(newRightSprings)

	return &newNode
}

func (sg SpringGroup) GeneratePossibilityTree() SpringGroupTreeNode {
	return *generateTreeNode(sg.springs)
}

func (sg SpringGroup) Arrangements() int {
	// Brute force method:
	// generate all possible variants of the spring group
	// eg, for each ? in the spring group, generate a variant with that ? replaced with a . and a variant with that ? replaced with a #

	// Then generate the associated group numbers for each variant
	// If it matches sg.groups, then it's a valid arrangement

	// This is basically a binary tree, where each node has two children, one for "." and one for "#"

	// this is gonna take a royal assload of time and memory

	// log.Println("Generating possibility tree for:")
	// log.Printf("      %v", sg.springs)
	possibilityTree := sg.GeneratePossibilityTree()

	// Find all the string values that don't contain "?"
	// These are the "leaves" of the tree
	// For each leaf, generate the associated group numbers
	// If it matches sg.groups, then it's a valid arrangement

	// Find all the leaves
	leaves := []string{}
	
	// This function returns void but mutates the leaves variable	
	var findLeaves func(node *SpringGroupTreeNode)
	findLeaves = func(node *SpringGroupTreeNode) {
		if node.left == nil && node.right == nil {
			leaves = append(leaves, node.springs)
		} else {
			if node.left != nil {
				findLeaves(node.left)
			}
			if node.right != nil {
				findLeaves(node.right)
			}
		}
	}

	findLeaves(&possibilityTree)

	// For each leaf, generate the associated group numbers and compare to sg.groups
	arrangements := 0

	for _, leaf := range leaves {
		leafGroups, err := getGroupNumbers(leaf)
		if err != nil {
			log.Println(color.Red + "ERROR: " + color.Reset + err.Error())
		}
		//log.Printf("Leaf: %s %v", leaf, leafGroups)

		if len(leafGroups) == len(sg.groups) {
			for idx, leafGroup := range leafGroups {
				if leafGroup != sg.groups[idx] {
					break
				}
				if leafGroup == sg.groups[idx] && idx == len(leafGroups) - 1 {
					arrangements++
				}
			}
		}
	}

	return arrangements
}
