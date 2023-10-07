package main

import (
	"fmt"
	"day1/util/fileReader"
	"strconv"
	"slices"
)

func firstPuzzle() {
	lines := fileReader.ReadLines("input.txt")

	var maxCalories int64 = 0
	var currentCalories int64 = 0
	for _, line := range lines {
		item, err := strconv.ParseInt(line, 10, 64)
		
		if err != nil {
			currentCalories = 0
			continue
		}

		currentCalories += item
		
		if currentCalories > maxCalories {
			maxCalories = currentCalories
		}
	}

	fmt.Printf("Max calories: %d\n", maxCalories) 
}

func searchHighest(bags []int64, exclude []int64) int64 {
	var highest int64 = 0
	for _, item := range bags {
		if item > highest && !slices.Contains(exclude, item) {
			highest = item
		}
	}

	return highest
}

func secondPuzzle() {
	lines := fileReader.ReadLines("input.txt")

	bags := []int64{}
	var currentCalories int64 = 0
	for _, line := range lines {
		item, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			bags = append(bags, currentCalories)
			currentCalories = 0
			continue
		}

		currentCalories += item
	}

	n1 := searchHighest(bags, []int64{})
	n2 := searchHighest(bags, []int64{n1})
	n3 := searchHighest(bags, []int64{n1, n2})

	fmt.Printf("Calories count: %d\n", n1 + n2 + n3)
}

func main() {
	firstPuzzle()
	secondPuzzle()
}
