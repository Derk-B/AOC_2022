package main

import (
	"fmt"
	"Day3/util/fileReader"
	"Day3/util/math"
)

func strContain(str string, c rune) bool {
	for _, a := range str {
		if a == c {
			return true
		}
	}

	return false
}

func processContent(line string) int {
	mid := len(line) / 2
	left := line[0 : mid]

	for _, c := range line[mid: len(line)] {
		if strContain(left, c) {
			return math.Mod(int(c - 38), 58)
		}
	}

	return 0
}

func firstPart() {
	lines := fileReader.ReadLines("input.txt")

	score := 0
	for _, line := range lines {
		score += processContent(line)
	}

	fmt.Println(score)
}

type BadgeFlags struct {
	a1 bool
	a2 bool
}

func findBadge(bags []string) int {
	badgeFlags := ['z'+1]BadgeFlags {}

	for _, c := range bags[0] {
		badgeFlags[c].a1 = true
	}

	for _, c := range bags[1] {
		badgeFlags[c].a2 = true
	}

	for _, c := range bags[2] {
		if badgeFlags[c].a1 && badgeFlags[c].a2 {
			return math.Mod(int(c - 38), 58)
		}
	}

	return 0
}

func secondPart() {
	lines := fileReader.ReadLines("input.txt")
	
	score := 0

	for i := 0; i < len(lines); i += 3 {
		score += findBadge(lines[i : i + 3])
	}

	fmt.Println(score)
}

func main() {
	firstPart()
	secondPart()
}
