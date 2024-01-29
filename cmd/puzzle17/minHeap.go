package main

import (
	"errors"
	"fmt"
	"log"
)

type MinheapNode struct {
	idx  int
	dist int
}

func (n MinheapNode) String() string {
	return "{idx: " + fmt.Sprint(n.idx) + ", dist: " + fmt.Sprint(n.dist) + "}"
}

type Minheap struct {
	length int
	data   []MinheapNode // The data is the index of the node
}

func NewMinheap() Minheap {
	return Minheap{
		length: 0,
		data:   []MinheapNode{},
	}
}

func (h *Minheap) Insert(idx int, dist int) {
	h.data = append(h.data, MinheapNode{idx: idx, dist: dist})
	h.length++
	h.bubbleUp(h.length - 1)
}

func (h *Minheap) Pop() (out MinheapNode, err error) {
	if h.length == 0 {
		return MinheapNode{idx: -1, dist: -1}, errors.New("Heap is empty")
	}

	out = h.data[0]

	if h.length == 1 {
		h.data = []MinheapNode{}
		h.length = 0
		return out, nil
	}

	lastElement := h.data[len(h.data)-1]

	// Take the last element
	h.data[0] = lastElement
	h.data = h.data[:len(h.data)-1]
	h.length--
	h.bubbleDown(0)

	return out, nil
}

func (h *Minheap) Update(idx int, newDist int) {
	var node *MinheapNode // The node we need to update
	var heapIndex int     // The position in the heap array where the target node is

	// Linear search for now I think
	for i := 0; i < h.length; i++ {
		n := h.data[i]

		if n.idx == idx {
			heapIndex = i
			node = &n
			break
		}
	}

	if node == nil { // aww yiss, pointer magic
		log.Printf("Node not found")
		return
	}

	// Update the node
	oldDist := node.dist
	node.dist = newDist

	// For some reason the address changes so we explicitly need to change it
	// I was hoping passing a pointer to the array element would be enough
	// but no dice
	h.data[heapIndex] = *node

	// Bubble
	if oldDist < newDist {
		h.bubbleDown(heapIndex)
	} else if oldDist > newDist {
		h.bubbleUp(heapIndex)
	}
}

func (h *Minheap) parent(index int) int {
	return (index - 1) / 2
}

func (h *Minheap) leftChild(index int) int {
	return index*2 + 1
}

func (h *Minheap) rightChild(index int) int {
	return index*2 + 2
}

func (h *Minheap) swap(index1 int, index2 int) {
	h.data[index1], h.data[index2] = h.data[index2], h.data[index1]
}

func (h *Minheap) bubbleUp(index int) {
	if index == 0 {
		return
	}

	value := h.data[index].dist
	parent := h.parent(index)
	parentValue := h.data[parent].dist

	if value < parentValue {
		h.swap(index, parent)
		h.bubbleUp(parent)
	}
}

func (h *Minheap) bubbleDown(index int) {
	if index >= h.length {
		return
	}

	lIdx := h.leftChild(index)
	rIdx := h.rightChild(index)

	// if lIdx exceeds length, then we have no children
	if lIdx >= h.length {
		return
	}

	lValue := h.data[lIdx].dist
	value := h.data[index].dist

	if rIdx >= h.length {
		// We have no right child, but we do have a left child
		// Swap if the left child is smaller

		if value > lValue {
			h.swap(index, lIdx)
			h.bubbleDown(lIdx)
		}

		return
	}

	// We have both children
	rValue := h.data[rIdx].dist

	if lValue > rValue && value > rValue {
		h.swap(index, rIdx)
		h.bubbleDown(rIdx)
	} else if lValue < rValue && value > lValue {
		h.swap(index, lIdx)
		h.bubbleDown(lIdx)
	}
}
