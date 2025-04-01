package main

import (
	"Day7/util/fileReader"
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	name     string
	parent   *Directory
	children map[string]*Directory
	files    map[string]int
	size     int
}

func calculateSize(dir *Directory) {
	size := 0
	for _, child := range dir.children {
		calculateSize(child)
		size += child.size
	}
	for _, fileSize := range dir.files {
		size += fileSize
	}
	dir.size = size
}

func sizeOfAllSmallDirectories(dir *Directory) int {
	total := 0
	if dir.size <= 100000 {
		total += dir.size
	}

	for _, child := range dir.children {
		total += sizeOfAllSmallDirectories(child)
	}

	return total
}

func sizeOfSmallestSufficientDirectory(dir *Directory, requiredSize int) int {
	size := 0
	if dir.size >= requiredSize {
		size = dir.size
	}

	for _, child := range dir.children {
		childSize := sizeOfSmallestSufficientDirectory(child, requiredSize)
		if childSize >= requiredSize && (childSize < size || size == 0) {
			size = childSize
		}
	}

	return size
}

func main() {
	lines := fileReader.ReadLines("input.txt")

	root := Directory{
		name:     "/",
		parent:   nil,
		children: make(map[string]*Directory),
		files:    make(map[string]int),
	}

	var current *Directory = nil

	for _, line := range lines {
		if line == "$ cd /" {
			current = &root
			continue
		}
		if line == "$ cd .." {
			current = current.parent
			continue
		}
		parts := strings.Split(line, " ")

		if parts[0] == "$" && parts[1] == "cd" {
			// Create new subDirectory if it doesn't exist
			if current.children[parts[2]] == nil {
				current.children[parts[2]] = &Directory{
					name:     parts[2],
					parent:   current,
					children: make(map[string]*Directory),
					files:    make(map[string]int),
				}
			}
			current = current.children[parts[2]]
		} else if parts[0] == "$" && parts[1] == "ls" {
		} else if parts[0] == "dir" {
			// Create new subDirectory if it doesn't exist
			if current.children[parts[1]] == nil {
				current.children[parts[1]] = &Directory{
					name:     parts[1],
					parent:   current,
					children: make(map[string]*Directory),
					files:    make(map[string]int),
				}
			}
		} else {
			size, e := strconv.Atoi(parts[0])
			if e != nil {
				fmt.Println(e)
				panic("Cannot convert to int")
			}
			current.files[parts[1]] = size
		}
	}

	calculateSize(&root)

	fmt.Println("Answer 1:", sizeOfAllSmallDirectories(&root))
	fmt.Println("Answer 2:", sizeOfSmallestSufficientDirectory(&root, root.size-40000000))
}
