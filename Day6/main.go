package main

import "Day6/util/fileReader"

func strIsUnique(str string) bool {
	m := make(map[rune]bool)
	for _, c := range str {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}

func main() {
	lines := fileReader.ReadLines("input.txt")
	line := lines[0]

	for i := 3; i < len(line); i++ {
		if strIsUnique(line[i-3 : i+1]) {
			println("Answer 1", i+1)
			break
		}
	}

	for i := 13; i < len(line); i++ {
		if strIsUnique(line[i-13 : i+1]) {
			println("Answer 2", i+1)
			break
		}
	}
}
