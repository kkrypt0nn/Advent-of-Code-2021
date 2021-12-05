package main

import (
	"fmt"
	"strconv"

	"co.uk.krypt0n/aoc/helpers"
)

func main() {
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	lines := helpers.ReadFile("input-test.txt")
	previous, total := 0, 0
	for _, line := range lines {
		lineInt, _ := strconv.Atoi(line)
		if lineInt > previous && previous != 0 {
			total += 1
		}
		previous = lineInt
	}
	return total
}

func partTwo() int {
	lines := helpers.ReadFile("input-test.txt")
	previous, total := 0, 0
	for line := range lines {
		var list []int
		for _, linee := range lines[line : line+3] {
			lineInt, _ := strconv.Atoi(linee)
			if lineInt == 0 {
				break
			}
			list = append(list, lineInt)
		}
		sum := helpers.Sum(list)
		if sum > previous && previous != 0 {
			total += 1
		}
		previous = sum
	}
	return total
}
