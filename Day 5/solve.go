package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"co.uk.krypt0n/aoc/helpers"
)

type diagram struct {
	grid [1000][1000]int
}

type coord struct {
	xStart int
	yStart int
	xEnd   int
	yEnd   int
}

func main() {
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	lines := helpers.ReadFile("input-test.txt")
	var coords []coord
	for _, line := range lines {
		var newCoord coord
		rawCoords := strings.Split(line, " -> ")
		for i, coord := range rawCoords {
			xInt, _ := strconv.Atoi(strings.Split(coord, ",")[0])
			yInt, _ := strconv.Atoi(strings.Split(coord, ",")[1])
			if i == 0 {
				newCoord.xStart, newCoord.yStart = xInt, yInt
			} else {
				newCoord.xEnd, newCoord.yEnd = xInt, yInt
			}
		}
		coords = append(coords, newCoord)
	}
	diagram := diagram{}
	for _, coord := range coords {
		if coord.xStart == coord.xEnd {
			if coord.yStart > coord.yEnd {
				for i := coord.yStart; i >= coord.yEnd; i-- {
					num := diagram.grid[i][coord.xStart]
					diagram.grid[i][coord.xStart] = num + 1
				}
			} else {
				for i := coord.yStart; i <= coord.yEnd; i++ {
					num := diagram.grid[i][coord.xStart]
					diagram.grid[i][coord.xStart] = num + 1
				}
			}
		} else if coord.yStart == coord.yEnd {
			if coord.xStart > coord.xEnd {
				for i := coord.xStart; i >= coord.xEnd; i-- {
					num := diagram.grid[coord.yStart][i]
					diagram.grid[coord.yStart][i] = num + 1
				}
			} else {
				for i := coord.xStart; i <= coord.xEnd; i++ {
					num := diagram.grid[coord.yStart][i]
					diagram.grid[coord.yStart][i] = num + 1
				}
			}
		}
	}
	sum := 0
	for _, line := range diagram.grid {
		for _, num := range line {
			if num >= 2 {
				sum += 1
			}
		}
	}
	return sum
}

func partTwo() int {
	lines := helpers.ReadFile("input-test.txt")
	var coords []coord
	for _, line := range lines {
		var newCoord coord
		rawCoords := strings.Split(line, " -> ")
		for i, coord := range rawCoords {
			xInt, _ := strconv.Atoi(strings.Split(coord, ",")[0])
			yInt, _ := strconv.Atoi(strings.Split(coord, ",")[1])
			if i == 0 {
				newCoord.xStart, newCoord.yStart = xInt, yInt
			} else {
				newCoord.xEnd, newCoord.yEnd = xInt, yInt
			}
		}
		coords = append(coords, newCoord)
	}
	diagram := diagram{}
	for _, coord := range coords {
		if coord.xStart == coord.xEnd {
			if coord.yStart > coord.yEnd {
				for i := coord.yStart; i >= coord.yEnd; i-- {
					num := diagram.grid[i][coord.xStart]
					diagram.grid[i][coord.xStart] = num + 1
				}
			} else {
				for i := coord.yStart; i <= coord.yEnd; i++ {
					num := diagram.grid[i][coord.xStart]
					diagram.grid[i][coord.xStart] = num + 1
				}
			}
		} else if coord.yStart == coord.yEnd {
			if coord.xStart > coord.xEnd {
				for i := coord.xStart; i >= coord.xEnd; i-- {
					num := diagram.grid[coord.yStart][i]
					diagram.grid[coord.yStart][i] = num + 1
				}
			} else {
				for i := coord.xStart; i <= coord.xEnd; i++ {
					num := diagram.grid[coord.yStart][i]
					diagram.grid[coord.yStart][i] = num + 1
				}
			}
		} else {
			diff := math.Abs(float64(coord.xEnd-coord.xStart)) + 1
			startX := coord.xStart
			startY := coord.yStart
			if coord.xStart < coord.xEnd && coord.yStart < coord.yEnd {
				for i := 1; i <= int(diff); i++ {
					num := diagram.grid[startY][startX]
					diagram.grid[startY][startX] = num + 1
					startX++
					startY++
				}
			} else if coord.xStart > coord.xEnd && coord.yStart < coord.yEnd {
				for i := 1; i <= int(diff); i++ {
					num := diagram.grid[startY][startX]
					diagram.grid[startY][startX] = num + 1
					startX--
					startY++
				}
			} else if coord.yStart > coord.yEnd && coord.xStart < coord.xEnd {
				for i := 1; i <= int(diff); i++ {
					num := diagram.grid[startY][startX]
					diagram.grid[startY][startX] = num + 1
					startX++
					startY--
				}
			} else if coord.yStart > coord.yEnd && coord.xStart > coord.xEnd {
				for i := 1; i <= int(diff); i++ {
					num := diagram.grid[startY][startX]
					diagram.grid[startY][startX] = num + 1
					startX--
					startY--
				}
			}
		}
	}
	sum := 0
	for _, line := range diagram.grid {
		for _, num := range line {
			if num >= 2 {
				sum += 1
			}
		}
	}
	return sum
}
