package main

import (
	"Day10/util/fileReader"
	"fmt"
	"strconv"
	"strings"
)

var cpuCycle = 0
var cpuValue = 1

func incCycle() {
	cpuCycle++
	if (cpuCycle-20)%40 == 0 {
		fmt.Println(cpuCycle*cpuValue, cpuCycle, cpuValue)
	}
}

func main() {
	lines := fileReader.ReadLines("test.txt")

	for _, line := range lines {
		segments := strings.Split(line, " ")
		if segments[0] == "noop" {
			incCycle()
		} else if segments[0] == "addx" {
			value, e := strconv.Atoi(segments[1])
			if e != nil {
				panic(e)
			}
			incCycle()
			incCycle()
			cpuValue += value
		}
	}
}
