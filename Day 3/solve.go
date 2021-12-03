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
	gamma := ""
	epsilon := ""
	for i := 0; i < len(lines[5]); i++ {
		countZero := 0
		countOne := 0
		for _, line := range lines {
			switch string(line[i]) {
			case "0":
				countZero++
			case "1":
				countOne++
			}
		}
		if countOne > countZero {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaInt) * int(epsilonInt)
}

/*
This code is a disaster and I know it.
I will try to optimise it when (and if) I have time :)
*/
func part_two() int {
	lines := helpers.ReadFile("input-test.txt")
	list := lines
	var newList []string
	for i := 0; i < len(lines[5]); i++ {
		countOne := 0
		countZero := 0
		for _, item := range list {
			switch string(item[i]) {
			case "0":
				countZero++
			case "1":
				countOne++
			}
		}
		for _, item := range list {
			if string(item[i]) == "1" && countOne > countZero {
				newList = append(newList, item)
			}
			if string(item[i]) == "0" && countOne < countZero {
				newList = append(newList, item)
			}
			if countOne == countZero {
				if string(item[i]) == "1" {
					newList = append(newList, item)
				}
			}
		}
		if len(list) != 1 {
			list = newList
		}
		newList = []string{}
	}
	oxygen, _ := strconv.ParseInt(list[0], 2, 64)

	list = lines
	for i := 0; i < len(lines[5]); i++ {
		countOne := 0
		countZero := 0
		for _, item := range list {
			switch string(item[i]) {
			case "0":
				countZero++
			case "1":
				countOne++
			}
		}
		for _, item := range list {
			if string(item[i]) == "1" && countOne < countZero {
				newList = append(newList, item)
			}
			if string(item[i]) == "0" && countOne > countZero {
				newList = append(newList, item)
			}
			if countOne == countZero {
				if string(item[i]) == "0" {
					newList = append(newList, item)
				}
			}
		}
		if len(list) != 1 {
			list = newList
		}
		newList = []string{}
	}
	carbon, _ := strconv.ParseInt(list[0], 2, 64)

	return int(oxygen) * int(carbon)
}
