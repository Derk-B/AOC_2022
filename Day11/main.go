package main

import (
	"Day11/util/fileReader"
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []int
	operation       string
	testDivisor     int
	monkeyWhenTrue  int
	monkeyWhenFalse int
}

var bigMod = 1

func parseOperation(line string, old int) int {
	strs := strings.Split(line, " ")

	leftStr := strs[5]
	operatorStr := strs[6]
	rightStr := strs[7]

	left, right := 0, 0

	if leftStr == "old" {
		left = old
	} else {
		n, e := strconv.Atoi(leftStr)
		left = n
		if e != nil {
			panic(e)
		}
	}
	if rightStr == "old" {
		right = old
	} else {
		n, e := strconv.Atoi(rightStr)
		right = n
		if e != nil {
			panic(e)
		}
	}

	if operatorStr == "+" {
		return left + right
	} else {
		return left * right
	}
}

func parseMonkeys(lines []string) []Monkey {
	monkeys := []Monkey{}

	for mi := 0; mi < (len(lines)+1)/7; mi++ {
		newMonkey := Monkey{}

		newMonkey.Items = make([]int, 0)
		itemStrs := strings.Split(lines[mi*7+1], " ")[4:]

		// Parse items
		for _, itemStr := range itemStrs {
			item, err := strconv.Atoi(strings.Trim(itemStr, ","))
			if err != nil {
				panic(err)
			}
			bigMod *= item
			newMonkey.Items = append(newMonkey.Items, item)
		}

		// Parse operation
		newMonkey.operation = lines[mi*7+2]

		// Parse test divisor
		divisorStr := strings.Split(lines[mi*7+3], " ")[5]
		divisor, err := strconv.Atoi(divisorStr)
		if err != nil {
			panic(err)
		}
		newMonkey.testDivisor = divisor

		// Parse true and false targets
		trueMonkeyStr := strings.Split(lines[mi*7+4], " ")[9]
		n, err := strconv.Atoi(trueMonkeyStr)
		if err != nil {
			panic(err)
		}
		newMonkey.monkeyWhenTrue = n
		falseMonkeyStr := strings.Split(lines[mi*7+5], " ")[9]
		n, err = strconv.Atoi(falseMonkeyStr)
		if err != nil {
			panic(err)
		}
		newMonkey.monkeyWhenFalse = n

		monkeys = append(monkeys, newMonkey)
	}
	return monkeys
}

func (m *Monkey) operate(monkeys *[]Monkey) {
	for _, item := range m.Items {
		// fmt.Println("Processing item", item)
		newItemValue := parseOperation(m.operation, item)
		newItemValue = newItemValue / 3

		itemCount := 0
		for _, m1 := range *monkeys {
			itemCount += len(m1.Items)
		}
		fmt.Println(itemCount)
		// newItemValue = newItemValue % bigMod
		if newItemValue%m.testDivisor == 0 {
			(*monkeys)[m.monkeyWhenTrue].Items = append((*monkeys)[m.monkeyWhenTrue].Items, newItemValue)
		} else {
			(*monkeys)[m.monkeyWhenFalse].Items = append((*monkeys)[m.monkeyWhenFalse].Items, newItemValue)
		}
	}
}

func main() {
	lines := fileReader.ReadLines("test.txt")
	monkeys := parseMonkeys(lines)
	for i := 0; i < 20; i++ {
		// fmt.Println("Round: ", i)
		for _, monkey := range monkeys {
			// fmt.Println("Processing a monkey")
			monkey.operate(&monkeys)
			monkey.Items = []int{}
		}
	}

	fmt.Println(monkeys)
}
