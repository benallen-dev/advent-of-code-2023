package main

import (
	"fmt"
)

// ProcessQueue holds a list of position tuples: [position, previousPosition]
type PositionQueue []QueueItem

type QueueItem struct {
	position        Position
	previousPosition Position
}

var positionQueue PositionQueue

func (pq PositionQueue) pop() QueueItem {
	item := positionQueue[0]
	positionQueue = positionQueue[1:]
	
	return item
}

func (pq PositionQueue) isEmpty() bool {
	return len(positionQueue) == 0
}

func (pq PositionQueue) String() string {
	var str string

	for _, item := range positionQueue {
		str += fmt.Sprintf("%v, %v\n", item.position, item.previousPosition)
	}
	return str
}

func (pq PositionQueue) initWith(position Position, previousPosition Position) {	
	positionQueue = PositionQueue{QueueItem{position, previousPosition}}
}

func addToQueue(position Position, previousPosition Position) {
	positionQueue = append(positionQueue, QueueItem{position, previousPosition})
}

func queueLeft(position Position) {
	newPosition := Position{position[0], position[1] - 1}
	if !isOutOfBounds(newPosition) {
		addToQueue(newPosition, position)
	}
}

func queueRight(position Position) {
	newPosition := Position{position[0], position[1] + 1}
	if !isOutOfBounds(newPosition) {
		addToQueue(newPosition, position)
	}
}

func queueUp(position Position) {
	newPosition := Position{position[0] - 1, position[1]}
	if !isOutOfBounds(newPosition) {
		addToQueue(newPosition, position)
	}
}

func queueDown(position Position) {
	newPosition := Position{position[0] + 1, position[1]}
	if !isOutOfBounds(newPosition) {
		addToQueue(newPosition, position)
	}
}

func queuePosition(newPosition Position, position Position) {
	if isOutOfBounds(newPosition) {
		return
	}

	addToQueue(newPosition, position)
}
