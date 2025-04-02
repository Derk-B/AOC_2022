package main

import (
	"Day8/util/fileReader"
	"fmt"
	"strconv"
)

type Location struct {
	x int
	y int
}

func part1(lines []string) {
	visibleTreeMap := make(map[Location]bool)

	for i := 0; i < len(lines); i++ {
		// Left to right
		currentTallest := -1
		for j := 0; j < len(lines[0]); j++ {
			currentHeight, e := strconv.Atoi(string(lines[i][j]))
			if e != nil {
				panic(e)
			}
			if currentHeight > currentTallest {
				currentTallest = currentHeight
				visibleTreeMap[Location{x: j, y: i}] = true
			}
		}
		// Right to left
		currentTallest = -1
		for j := len(lines[0]) - 1; j >= 0; j-- {
			currentHeight, e := strconv.Atoi(string(lines[i][j]))
			if e != nil {
				panic(e)
			}
			if currentHeight > currentTallest {
				currentTallest = currentHeight
				visibleTreeMap[Location{x: j, y: i}] = true
			}
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		// Top to bottom
		currentTallest := -1
		for j := 0; j < len(lines); j++ {
			currentHeight, e := strconv.Atoi(string(lines[j][i]))
			if e != nil {
				panic(e)
			}
			if currentHeight > currentTallest {
				currentTallest = currentHeight
				visibleTreeMap[Location{x: i, y: j}] = true
			}
		}
		// Right to left
		currentTallest = -1
		for j := len(lines[0]) - 1; j >= 0; j-- {
			currentHeight, e := strconv.Atoi(string(lines[j][i]))
			if e != nil {
				panic(e)
			}
			if currentHeight > currentTallest {
				currentTallest = currentHeight
				visibleTreeMap[Location{x: i, y: j}] = true
			}
		}
	}

	println("Answer 1: ", len(visibleTreeMap))
}

func part2(lines []string) {
	maxSceneryScore := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			ownHeight, e := strconv.Atoi(string(lines[y][x]))
			if e != nil {
				panic(e)
			}
			// Left
			visibleLeft := 0
			for i := x - 1; i >= 0; i-- {
				heightLeft, e := strconv.Atoi(string(lines[y][i]))
				if e != nil {
					panic(e)
				}
				visibleLeft++
				if heightLeft >= ownHeight {
					break
				}
			}

			// Right
			visibleRight := 0
			for i := x + 1; i < len(lines[0]); i++ {
				heightRight, e := strconv.Atoi(string(lines[y][i]))
				if e != nil {
					panic(e)
				}
				visibleRight++
				if heightRight >= ownHeight {
					break
				}
			}

			// Up
			visibleUp := 0
			for i := y - 1; i >= 0; i-- {
				heightUp, e := strconv.Atoi(string(lines[i][x]))
				if e != nil {
					panic(e)
				}
				visibleUp++
				if heightUp >= ownHeight {
					break
				}
			}

			// Down
			visibleDown := 0
			for i := y + 1; i < len(lines); i++ {
				heightDown, e := strconv.Atoi(string(lines[i][x]))
				if e != nil {
					panic(e)
				}
				visibleDown++
				if heightDown >= ownHeight {
					break
				}
			}

			sceneryScore := visibleLeft * visibleRight * visibleUp * visibleDown
			maxSceneryScore = max(maxSceneryScore, sceneryScore)
		}
	}

	fmt.Println("Answer 2: ", maxSceneryScore)
}

func main() {
	lines := fileReader.ReadLines("input.txt")

	part1(lines)
	part2(lines)
}
