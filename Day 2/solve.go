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
	posX := 0
	posY := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		lineInt, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			posX += lineInt
		case "down":
			posY += lineInt
		case "up":
			posY -= lineInt
		}
	}
	return posX * posY
}

func partTwo() int {
	lines := helpers.ReadFile("input-test.txt")
	posX := 0
	posY := 0
	aim := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		lineInt, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			posX += lineInt
			posY += aim * lineInt
		case "down":
			aim += lineInt
		case "up":
			aim -= lineInt
		}
	}
	return posX * posY
}
