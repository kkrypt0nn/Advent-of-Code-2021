package main

import (
	"fmt"
	"strconv"
	"strings"

	"co.uk.krypt0n/aoc/helpers"
)

type diagram struct {
	grid   [][]int
	height int
	width  int
}

type coord struct {
	xStart int
	yStart int
	xEnd   int
	yEnd   int
}

func main() {
	fmt.Println("Part one:", partOne())
	// fmt.Println("Part two:", partTwo())
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
	xMax, yMax := 0, 0
	for _, coord := range coords {
		if coord.xEnd > xMax {
			xMax = coord.xEnd
		}
		if coord.xStart > xMax {
			xMax = coord.xStart
		}
		if coord.yEnd > yMax {
			yMax = coord.yEnd
		}
		if coord.yStart > yMax {
			yMax = coord.yStart
		}
	}
	diagram := diagram{}
	for i := 0; i <= xMax; i++ {
		diagram.grid = append(diagram.grid, make([]int, xMax+1))
	}
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
