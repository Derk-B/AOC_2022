package main

import (
	"Day9/util/fileReader"
	"fmt"
	"strconv"
)

type Location struct {
	x int
	y int
}

func dirToVec(dir rune) Location {
	switch dir {
	case 'U':
		return Location{0, 1}
	case 'D':
		return Location{0, -1}
	case 'L':
		return Location{-1, 0}
	case 'R':
		return Location{1, 0}
	default:
		panic("Invalid direction")
	}
}

func moveTail(t Location, h Location) Location {
	// UP
	if h.x == t.x && h.y-t.y > 1 {
		return Location{t.x, t.y + 1}
	} else
	// Down
	if h.x == t.x && t.y-h.y > 1 {
		return Location{t.x, t.y - 1}
	} else
	// Right
	if h.y == t.y && h.x-t.x > 1 {
		return Location{t.x + 1, t.y}
	} else
	// Left
	if h.y == t.y && t.x-h.x > 1 {
		return Location{t.x - 1, t.y}
	}

	return h
}

func main() {
	lines := fileReader.ReadLines("test.txt")
	visitedMap := make(map[Location]bool)

	head := Location{0, 0}
	tail := Location{0, 0}
	visitedMap[tail] = true
	for _, line := range lines {
		direction := line[0]
		steps, e := strconv.Atoi(line[2:])
		if e != nil {
			panic("Invalid step count")
		}
		vector := dirToVec(rune(direction))

		for i := 0; i < steps; i++ {
			head.x += vector.x
			head.y += vector.y

			tail := moveTail(tail, head)
			visitedMap[tail] = true
		}
	}

	fmt.Println(len(visitedMap))
	fmt.Println(visitedMap)
}
