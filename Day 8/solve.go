package main

import (
	"fmt"
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
	numbers := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}
	for _, line := range lines {
		split := strings.Split(line, " | ")
		values := strings.Split(split[1], " ")
		for _, value := range values {
			switch len(value) {
			case 2:
				numbers[1]++
			case 4:
				numbers[4]++
			case 3:
				numbers[7]++
			case 7:
				numbers[8]++
			}
		}
	}
	return helpers.MapSum(numbers)
}

func partTwo() int {
	lines := helpers.ReadFile("input-test.txt")
	total := 0
	for _, line := range lines {
		numbers := []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}
		known := map[int]string{
			1: "",
			4: "",
			7: "",
			8: "",
			// -- Unknown --
			// Len 6
			0: "",
			6: "",
			9: "",
			// Len 5
			2: "",
			3: "",
			5: "",
		}
		unknown := []string{}
		split := strings.Split(line, " | ")
		values := strings.Split(split[0], " ")
		for _, value := range values {
			switch len(value) {
			case 2:
				known[1] = string(value)
				numbers[1] = string(value)
			case 4:
				known[4] = string(value)
				numbers[4] = string(value)
			case 3:
				known[7] = string(value)
				numbers[7] = string(value)
			case 7:
				known[8] = string(value)
				numbers[8] = string(value)
			default:
				unknown = append(unknown, string(value))
			}
		}
		for _, value := range unknown {
			if len(value) == 6 {
				if helpers.Contains(value, known[4]) {
					// 9 has len 6 and only one to contain all in 4
					known[9] = string(value)
				} else if helpers.Contains(value, known[7]) && !helpers.Contains(value, known[4]) {
					// 0 has len 6 and only one to contain all in 7
					known[0] = string(value)
				} else {
					known[6] = string(value)
				}
			}
		}
		for _, value := range unknown {
			if len(value) == 5 {
				if helpers.Contains(value, known[1]) {
					// 3 has len 5 and only one to contain all in 1
					known[3] = string(value)
				} else if helpers.Contains(known[6], value) {
					// 5 has len 5 and only one to be fully contained in 6
					known[5] = string(value)
				} else {
					known[2] = string(value)
				}
			}
		}
		outputs := strings.Split(split[1], " ")
		value := ""
		for _, output := range outputs {
			for i, number := range known {
				if helpers.ContainsExactly(output, number) {
					value += strconv.Itoa(i)
				}
			}
		}
		outputInt, _ := strconv.Atoi(value)
		total += outputInt
	}
	return total
}
