package main

import (
	"Day9/util/fileReader"
	"fmt"
	"math"
	"strconv"
)

type Location struct {
	x int
	y int
}

type Knot struct {
	pos  Location
	prev Location
	name string
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

func round(x float64) int {
	if x > 0 {
		return int(math.Ceil(x))
	} else {
		return int(math.Floor(x))
	}
}

func moveTail(t Location, h Location, hPrev Location) Location {
	dx := h.x - t.x
	dy := h.y - t.y
	absDX := math.Abs(float64(dx))
	absDY := math.Abs(float64(dy))
	if absDX > 1 || absDY > 1 {
		return Location{t.x + round(float64(dx)/2), t.y + round(float64(dy)/2)}
	} else {
		return t
	}
}

func print(knots [10]Knot) {
	for y := -13; y < 7; y++ {
		for x := -13; x < 15; x++ {
			found := false
			for _, knot := range knots {
				if knot.pos.x == x && knot.pos.y == -y {
					fmt.Print(knot.name)
					found = true
					break
				}
			}
			if found == false {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part1() {
	lines := fileReader.ReadLines("input.txt")
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
			prevHead := head
			head.x += vector.x
			head.y += vector.y

			nextPos := moveTail(tail, head, prevHead)
			tail.x = nextPos.x
			tail.y = nextPos.y
			visitedMap[tail] = true
		}
	}

	fmt.Println("Answer 1: ", len(visitedMap))
}

func part2() {
	knots := [10]Knot{
		{Location{0, 0}, Location{0, 0}, "H"},
		{Location{0, 0}, Location{0, 0}, "1"},
		{Location{0, 0}, Location{0, 0}, "2"},
		{Location{0, 0}, Location{0, 0}, "3"},
		{Location{0, 0}, Location{0, 0}, "4"},
		{Location{0, 0}, Location{0, 0}, "5"},
		{Location{0, 0}, Location{0, 0}, "6"},
		{Location{0, 0}, Location{0, 0}, "7"},
		{Location{0, 0}, Location{0, 0}, "8"},
		{Location{0, 0}, Location{0, 0}, "T"},
	}
	visitedMap := make(map[Location]bool)
	visitedMap[knots[len(knots)-1].pos] = true

	headVisited := make(map[Location]bool)
	headVisited[knots[0].pos] = true

	for _, line := range fileReader.ReadLines("input.txt") {
		direction := line[0]
		steps, e := strconv.Atoi(line[2:])
		if e != nil {
			panic("Invalid step count")
		}
		vector := dirToVec(rune(direction))
		for i := 0; i < steps; i++ {
			knots[0].prev.x = knots[0].pos.x
			knots[0].prev.y = knots[0].pos.y
			knots[0].pos.x += vector.x
			knots[0].pos.y += vector.y
			headVisited[knots[0].pos] = true

			for j := 1; j < len(knots); j++ {
				nextPos := moveTail(knots[j].pos, knots[j-1].pos, knots[j-1].prev)
				knots[j].prev.x = knots[j].pos.x
				knots[j].prev.y = knots[j].pos.y
				knots[j].pos.x = nextPos.x
				knots[j].pos.y = nextPos.y

				if j == len(knots)-1 {
					visitedMap[knots[j].pos] = true
				}
			}
		}
	}

	fmt.Println("Answer 2: ", len(visitedMap))
}

func main() {
	part1()
	part2()
}
