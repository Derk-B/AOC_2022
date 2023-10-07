package main

import (
	"fmt"
	"Day2/util/fileReader"
)

// A function to replace the % operator.
// Go % is a remainder operator and not a modulo.
func mod(a int, n int) int {
	return ((a % n) + n) % n
}

func firstPart() {
	lines := fileReader.ReadLines("input.txt")

	score := 0
	for _, line := range lines {
		player := int(line[2]) - 87
		other := int(line[0]) - 64

		// Add shape score
		score += player

		// Add game score
		score += mod(mod(player - other, 3) * 3 + 3, 9)
	}
	
	fmt.Println(score)
}

func secondPart() {
	lines := fileReader.ReadLines("input.txt")

	score := 0
	for _, line := range lines {
		goal := int(line[2]) - 87
		other := int(line[0]) - 64

		// Add shape score
		score += mod(other - (3 - goal), 3) + 1
		
		// Add game score
		score += (goal - 1) * 3
	}

	fmt.Println(score)
}

func main() {
	firstPart()
	secondPart()
}
