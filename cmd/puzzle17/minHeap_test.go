package main

import (
	"testing"
)

func TestMinheap(t *testing.T) {
	minheap := NewMinheap()

	// Let's insert some values, idx not important but dist is
	minheap.Insert(0, 1)	
	minheap.Insert(1, 2)
	minheap.Insert(2, 9)
	minheap.Insert(3, 6)
	minheap.Insert(4, 5)
	minheap.Insert(5, 3)
	minheap.Insert(6, 4)
	minheap.Insert(7, 8)
	minheap.Insert(8, 7)

	t.Logf("Minheap: %v", minheap)

	for i := 1; i < 10; i++ {
		node, err := minheap.Pop()

		if err != nil {
			t.Errorf("Error popping minheap: %s", err)
		}

		if node.dist != i {
			t.Errorf("Expected %d, got %d", i, node.dist)
		}
	}
}
