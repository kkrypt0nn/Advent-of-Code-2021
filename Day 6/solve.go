package main

import (
	"fmt"
	"strconv"
	"strings"

	"co.uk.krypt0n/aoc/helpers"
)

func main() {
	fmt.Println("Part one:", solve(80))
	fmt.Println("Part two:", solve(256))
}

func solve(days int) int {
	lines := helpers.ReadFile("input-test.txt")
	values := strings.Split(lines[0], ",")
	valuesInt := []int{}
	for i := range values {
		valueInt, _ := strconv.Atoi(values[i])
		valuesInt = append(valuesInt, valueInt)
	}
	// No bruteforce, saving number of fishes at each day
	fishes := [9]int{}
	for _, x := range valuesInt {
		fishes[x] += 1
	}
	for i := 0; i < days; i++ {
		// Get number of fishes at 0
		totalAtZero := fishes[0]
		// Move all numbers of fishes to the left
		for j := 1; j < len(fishes); j++ {
			fishes[j-1] = fishes[j]
		}
		// Add the number of fishes at 0 to the number of fishes at reset (6)
		fishes[6] += totalAtZero
		// Fishes at max (8) are set to number of previous fishes at 0
		fishes[8] = totalAtZero
	}
	return helpers.Sum(fishes[:])
}
