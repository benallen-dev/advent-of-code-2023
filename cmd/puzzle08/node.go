package main

import (
	"fmt"
)

type Node struct {
	name string
	left string
	right string
}

func (n Node) String() string {
	return fmt.Sprintf("%s = (%s, %s)", n.name, n.left, n.right)
}

