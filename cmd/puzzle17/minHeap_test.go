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

func TestUpdate(t *testing.T) {
	minheap := NewMinheap()

	// Let's insert some values, idx not important but dist is
	minheap.Insert(9, 69420)
	minheap.Insert(0, 10)	
	minheap.Insert(1, 2)
	minheap.Insert(2, 9)
	minheap.Insert(3, 6)
	minheap.Insert(4, 5)
	minheap.Insert(5, 3)
	minheap.Insert(6, 4)
	minheap.Insert(7, 8)
	minheap.Insert(8, 7)

	t.Logf("Minheap:                   %v", minheap)

	minheap.Update(8, 1)
	t.Logf("Minheap after Update(8,1): %v", minheap)

	// We now expect
	// [{1 2} {8 1} {5 3} {6 4} {4 5} {3 6} {2 9} {7 8} {0 10}] - ish
	// pop should return 1

	min, err := minheap.Pop()

	if err != nil {
		t.Errorf("Error popping minheap: %s", err)
	}

	t.Logf("min: %v", min)

	if min.dist != 1 {
		t.Errorf("Expected 1, got %d", min.dist)
	}


}
