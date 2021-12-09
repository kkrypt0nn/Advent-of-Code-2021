package main

import (
	"fmt"
	"strconv"

	"co.uk.krypt0n/aoc/helpers"
)

func main() {
	fmt.Println("Part one:", partOne())
}

func partOne() int {
	lines := helpers.ReadFile("input-test.txt")
	list := []int{}
	for rowIndex, row := range lines {
		for charIndex := range row {
			left := 99
			right := 99
			up := 99
			down := 99
			if rowIndex != 0 {
				up, _ = strconv.Atoi(string(lines[rowIndex-1][charIndex]))
			}
			if rowIndex != len(lines)-1 {
				down, _ = strconv.Atoi(string(lines[rowIndex+1][charIndex]))
			}
			if charIndex != 0 {
				left, _ = strconv.Atoi(string(row[charIndex-1]))
			}
			if charIndex != len(row)-1 {
				right, _ = strconv.Atoi(string(row[charIndex+1]))
			}

			char, _ := strconv.Atoi(string(row[charIndex]))

			if char < left && char < right && char < up && char < down {
				list = append(list, char+1)
			}
		}
	}
	return helpers.Sum(list)
}
