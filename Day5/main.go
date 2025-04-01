package main

import (
	"Day5/util/fileReader"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// My challenge input contains exactly 9 stacks.
const STACK_COUNT = 9

type stack []byte

func (s *stack) Push(v byte) {
	*s = append((*s), v)
}

func (s *stack) Pop() (byte, error) {
	len := len(*s)

	if len == 0 {
		return 0, errors.New("cannot pop from empty stack")
	}

	val := (*s)[len-1]
	*s = (*s)[:len-1]
	return val, nil
}

func splitStacksAndInstructions(lines []string) ([]string, []string) {
	var stacks = []string{}
	instructions := []string{}

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			instructions = lines[i+1:]
			break
		}
		stacks = append(stacks, lines[i])
	}

	return stacks, instructions
}

func part1(stacks [STACK_COUNT]stack, instructions []string) {
	for i := 0; i < len(instructions); i++ {
		pieces := strings.Split(instructions[i], " ")
		amount, e1 := strconv.Atoi(pieces[1])
		origin, e2 := strconv.Atoi(pieces[3])
		target, e3 := strconv.Atoi(pieces[5])

		if e1 != nil || e2 != nil || e3 != nil {
			panic("Failed to parse instruction")
		}

		for j := 0; j < amount; j++ {
			v, e := stacks[origin-1].Pop()
			if e != nil {
				panic("Tried to pop from an empty stack")
			}
			stacks[target-1].Push(v)
		}
	}

	answer := ""
	for i := 0; i < len(stacks); i++ {
		c, err := stacks[i].Pop()
		if err != nil {
			panic("Failed to pop from result stack")
		}
		answer += string(c)
	}
	fmt.Println("Answer 1: ", answer)
}

func part2(stacks [STACK_COUNT]stack, instructions []string) {
	for i := 0; i < len(instructions); i++ {
		// Parse instruction
		pieces := strings.Split(instructions[i], " ")
		amount, e1 := strconv.Atoi(pieces[1])
		origin, e2 := strconv.Atoi(pieces[3])
		target, e3 := strconv.Atoi(pieces[5])

		if e1 != nil || e2 != nil || e3 != nil {
			panic("Failed to parse instruction")
		}

		// Move elements between stacks
		bytesToMove := []byte{}
		for j := 0; j < amount; j++ {
			v, e := stacks[origin-1].Pop()
			if e != nil {
				panic("Tried to pop from an empty stack")
			}
			bytesToMove = append(bytesToMove, v)
		}
		for j := len(bytesToMove) - 1; j >= 0; j-- {
			stacks[target-1].Push(bytesToMove[j])
		}
	}

	answer := ""
	for i := 0; i < len(stacks); i++ {
		c, err := stacks[i].Pop()
		if err != nil {
			panic("Failed to pop from result stack")
		}
		answer += string(c)
	}
	fmt.Println("Answer 2: ", answer)
}

func main() {
	lines := fileReader.ReadLines("input.txt")
	stacksStrs, instructions := splitStacksAndInstructions(lines)
	stacksP1 := [STACK_COUNT]stack{}
	stacksP2 := [STACK_COUNT]stack{}

	// Load stacks
	for i := len(stacksStrs) - 2; i >= 0; i-- {
		for x := 0; x < STACK_COUNT; x++ {
			pos := 4*x + 1
			c := stacksStrs[i][pos]
			if c == ' ' {
				continue
			}
			stacksP1[x].Push(c)
			stacksP2[x].Push(c)
		}
	}

	part1(stacksP1, instructions)
	part2(stacksP2, instructions)
}
