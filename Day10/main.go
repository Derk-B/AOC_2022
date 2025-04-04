package main

import (
	"Day10/util/fileReader"
	"fmt"
	"strconv"
	"strings"
)

var part1Res = 0
var part2Str = ""

func incCycle(cycle *int, value int) {
	(*cycle)++

	if *cycle%40 >= value && *cycle%40 < value+3 {
		part2Str += "#"
	} else {
		part2Str += "."
	}

	if (*cycle-20)%40 == 0 {
		part1Res += *cycle * value
	}
}

func main() {
	lines := fileReader.ReadLines("input.txt")

	cpuCycle := 0
	cpuValue := 1

	for _, line := range lines {
		segments := strings.Split(line, " ")
		if segments[0] == "noop" {
			incCycle(&cpuCycle, cpuValue)
		} else if segments[0] == "addx" {
			value, e := strconv.Atoi(segments[1])
			if e != nil {
				panic(e)
			}
			incCycle(&cpuCycle, cpuValue)
			incCycle(&cpuCycle, cpuValue)
			cpuValue += value
		}
	}

	fmt.Println("Part 1 Result: ", part1Res)
	fmt.Println("Part 2 Result: ")
	for i := 0; i < len(part2Str); i++ {
		if i%40 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Print(string(part2Str[i]))
	}
	fmt.Println()
}
