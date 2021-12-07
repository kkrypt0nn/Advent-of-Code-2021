package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"co.uk.krypt0n/aoc/helpers"
)

func main() {
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	lines := helpers.ReadFile("input-test.txt")
	values := strings.Split(lines[0], ",")
	valuesInt := []int{}
	for i := range values {
		valueInt, _ := strconv.Atoi(values[i])
		valuesInt = append(valuesInt, valueInt)
	}
	sort.Ints(valuesInt)
	median := valuesInt[len(valuesInt)/2]
	if len(valuesInt)%2 == 0 {
		median = (valuesInt[len(valuesInt)/2] + valuesInt[len(valuesInt)/2-1]) / 2
	}
	fuel := 0
	for _, value := range valuesInt {
		used := math.Abs(float64(value - median))
		fuel += int(used)
	}
	return fuel
}

func partTwo() int {
	lines := helpers.ReadFile("input-test.txt")
	values := strings.Split(lines[0], ",")
	valuesInt := []int{}
	for i := range values {
		valueInt, _ := strconv.Atoi(values[i])
		valuesInt = append(valuesInt, valueInt)
	}
	max := 1000000000 // High integer needed to make sure the result requires less fuel than that
	for i := 0; i < helpers.Max(valuesInt); i++ {
		fuel := []int{}
		for _, value := range valuesInt {
			prevFuel := 0
			used := math.Abs(float64(value - i))
			for j := int(used); j > 0; j-- {
				prevFuel += j
			}
			fuel = append(fuel, prevFuel)
		}
		if helpers.Sum(fuel) < max {
			max = helpers.Sum(fuel)
		}
	}
	return max
}
