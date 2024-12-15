package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func solve(fileName string) int {
	var warehouse [][]byte
	var res int
	var movesIdx int
	var row, col, dr, dc int
	var canMove func(int, int) bool
	input := utils.ReadMatrix(fileName)

	for i, line := range input {
		if len(line) == 0 {
			movesIdx = i
			break
		}
		warehouse = append(warehouse, line)
		if j := strings.Index(string(line), "@"); j != -1 {
			row, col = i, j
		}
	}

	var moves []byte
	for _, line := range input[movesIdx+1:] {
		moves = append(moves, line...)
	}

	canMove = func(r, c int) bool {
		if warehouse[r][c] == '.' {
			return true
		}
		if warehouse[r][c] == '#' {
			return false
		}

		if canMove(r+dr, c+dc) {
			cell := warehouse[r][c]
			warehouse[r+dr][c+dc] = cell
			warehouse[r][c] = '.'
			return true
		}
		return false
	}

	for _, move := range moves {
		dr, dc = 0, 0

		switch move {
		case '^':
			dr = -1
		case 'v':
			dr = 1
		case '<':
			dc = -1
		case '>':
			dc = 1
		}
		if canMove(row, col) {
			row += dr
			col += dc
		}
	}

	for i, row := range warehouse {
		for j, cell := range row {
			if cell == 'O' {
				res += i*100 + j
			}
		}
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 2028},
		{"../test2.txt", 10092},
		{"../input.txt", 1441031},
	}

	for _, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
