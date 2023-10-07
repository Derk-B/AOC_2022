package main

import (
	"fmt"
	"Day4/util/fileReader"
	"strings"
	"strconv"
)

func firstPart() {
	lines := fileReader.ReadLines("input.txt")

	
	score1 := 0
	score2 := 0
	for _, l := range lines {
		parts := strings.Split(l, ",")
		left := strings.Split(parts[0], "-")
		right := strings.Split(parts[1], "-")

		lStart, _ := strconv.Atoi(left[0])
		rStart, _ := strconv.Atoi(right[0])
		lEnd, _ := strconv.Atoi(left[1])
		rEnd, _ := strconv.Atoi(right[1])

		if lStart >= rStart && lEnd <= rEnd || lStart <= rStart && lEnd >= rEnd {
			score1++
		}

		if (lStart >= rStart && lStart <= rEnd) || (lEnd >= rStart && lEnd <= rEnd) || (rStart >= lStart && rStart <= lEnd) || (rEnd >= lStart && rEnd <= lEnd) {
			score2++
		}
	}

	fmt.Println(score1)
	fmt.Println(score2)
}

func main() {
	firstPart()
}
