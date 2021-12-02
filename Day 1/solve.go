package main

import (
	"fmt"
	"strconv"

	"co.uk.krypt0n/aoc/helpers"
)

func main() {
	fmt.Println("Part one:", part_one())
	fmt.Println("Part two:", part_two())
}

func part_one() int {
	lines := helpers.ReadFile("input-test.txt")
	previous := 0
	total := 0
	for _, line := range lines {
		lineInt, _ := strconv.Atoi(line)
		if lineInt > previous && previous != 0 {
			total += 1
		}
		previous = lineInt
	}
	return total
}

func part_two() int {
	lines := helpers.ReadFile("input-test.txt")
	previous := 0
	total := 0
	for line, _ := range lines {
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
