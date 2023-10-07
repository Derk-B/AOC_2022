package main

import (
	"fmt"
	"Day4/util/fileReader"
	"strings"
	"strconv"
)

func firstPart() {
	lines := fileReader.ReadLines("input.txt")

	
	score := 0
	for _, l := range lines {
		parts := strings.Split(l, ",")
		left := strings.Split(parts[0], "-")
		right := strings.Split(parts[1], "-")

		lStart, _ := strconv.Atoi(left[0])
		rStart, _ := strconv.Atoi(right[0])
		lEnd, _ := strconv.Atoi(left[1])
		rEnd, _ := strconv.Atoi(right[1])

		if lStart >= rStart && lEnd <= rEnd || lStart <= rStart && lEnd >= rEnd {
			score ++
		}
	}

	fmt.Println(score)
}

func main() {
	firstPart()
}
