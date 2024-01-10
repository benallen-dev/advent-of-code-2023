package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/benallen-dev/advent-of-code-2023/pkg/color"
)

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

func (sg SpringGroup) GeneratePossibilityTree() (SpringGroupTreeNode, error) {
	rootNode := generateTreeNode(sg.springs, sg.groups)
	if rootNode == nil {
		log.Println(color.Red + "ERROR: " + color.Reset + "sg.springs: " + sg.springs)
		log.Println(color.Red + "ERROR: " + color.Reset + "sg.groups: " + fmt.Sprintf("%v", sg.groups))

		return SpringGroupTreeNode{}, fmt.Errorf("rootNode is nil")
	}

	return *rootNode, nil
	//return *generateTreeNode(sg.springs, sg.groups)
}

func (sg SpringGroup) Arrangements() int {
	// generate all possible arrangements of springs
	possibilityTree, err := sg.GeneratePossibilityTree()
	if err != nil {
		log.Fatal(color.Red + "ERROR: " + color.Reset + err.Error())
	}

	// Find all the leaves
	leaves := []string{}
	
	// This function returns void but mutates the leaves variable	
	// It's a function because recursion
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

	arrangements := 0

	// For each leaf, generate the associated group numbers and compare to sg.groups
	for _, leaf := range leaves {
		if strings.Contains(leaf, "?") {
			// Not a legit leaf but an optimised branch that was skipped
			continue
		}

		leafGroups, err := getGroupNumbers(leaf)
		if err != nil {
			log.Println(color.Red + "ERROR: " + color.Reset + err.Error())
		}

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
