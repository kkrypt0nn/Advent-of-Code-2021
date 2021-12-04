package main

import (
	"fmt"
	"strconv"
	"strings"

	"co.uk.krypt0n/aoc/helpers"
)

const (
	rows = 5
	cols = 5
)

type bingoBoard struct {
	board [][]int
	draws [rows][cols]bool
}

func (b *bingoBoard) check(num int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.board[i][j] == num {
				b.draws[i][j] = true
			}
		}
	}
}

func (b *bingoBoard) checkWin() bool {
	for i := 0; i < rows; i++ {
		if b.draws[i][0] && b.draws[i][1] && b.draws[i][2] && b.draws[i][3] && b.draws[i][4] {
			return true
		}
	}
	for i := 0; i < cols; i++ {
		if b.draws[0][i] && b.draws[1][i] && b.draws[2][i] && b.draws[3][i] && b.draws[4][i] {
			return true
		}
	}
	return false
}

func (b *bingoBoard) getSumUnmarked() int {
	sum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !b.draws[i][j] {
				sum += b.board[i][j]
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

func partOne() int {
	lines := helpers.ReadFile("input-test.txt")
	draws := strings.Split(lines[0], ",")
	boards := getBoards(lines)
	for _, draw := range draws {
		drawInt, _ := strconv.Atoi(draw)
		for i := 0; i < len(boards); i++ {
			boards[i].check(drawInt)
			if boards[i].checkWin() {
				return drawInt * boards[i].getSumUnmarked()
			}
		}
	}
	return 0
}

func partTwo() int {
	lines := helpers.ReadFile("input-test.txt")
	draws := strings.Split(lines[0], ",")
	boards := getBoards(lines)
	hasPreviouslyWon := false
	for draw := range draws {
		if hasPreviouslyWon {
			// For some reason, if a board has won with a draw, it won't check the draw on other boards, therefore we need to check all boards again after a win
			hasPreviouslyWon = false
			drawInt, _ := strconv.Atoi(draws[draw-1])
			for i := 0; i < len(boards); i++ {
				boards[i].check(drawInt)
				if boards[i].checkWin() {
					if len(boards) == 1 {
						return drawInt * boards[i].getSumUnmarked()
					} else {
						boards = append(boards[:i], boards[i+1:]...)
						hasPreviouslyWon = true
					}
				}
			}
		}
		drawInt, _ := strconv.Atoi(draws[draw])
		for i := 0; i < len(boards); i++ {
			boards[i].check(drawInt)
			if boards[i].checkWin() {
				if len(boards) == 1 {
					return drawInt * boards[i].getSumUnmarked()
				} else {
					boards = append(boards[:i], boards[i+1:]...)
					hasPreviouslyWon = true
				}
			}
		}
	}
	return 0
}

func getBoards(lines []string) []bingoBoard {
	newBoard := false
	boards := []bingoBoard{}
	for x := 1; x < len(lines); x++ {
		if lines[x] == "" {
			newBoard = true
			continue
		}
		if newBoard {
			board := bingoBoard{}
			for i := x; i < x+rows; i++ {
				list := []int{}
				for _, v := range strings.Fields(lines[i]) {
					num, _ := strconv.Atoi(v)
					list = append(list, num)
				}
				board.board = append(board.board, list)
			}
			boards = append(boards, board)
			x += cols - 1
		}
	}
	return boards
}
