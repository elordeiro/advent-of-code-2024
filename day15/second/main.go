package main

import (
	"adventofcode/utils"
	"fmt"
)

func solve(fileName string) int {
	var warehouse [][]byte
	var res int
	var movesIdx int
	var row, col, dr, dc int
	var canMove func(int, int, bool) bool
	input := utils.ReadMatrix(fileName)

	for i, line := range input {
		if len(line) == 0 {
			movesIdx = i
			break
		}
		var newLine []byte
		for _, tile := range line {
			switch tile {
			case '@':
				row, col = i, len(newLine)
				newLine = append(newLine, '@', '.')
			case '#':
				newLine = append(newLine, '#', '#')
			case 'O':
				newLine = append(newLine, '[', ']')
			case '.':
				newLine = append(newLine, '.', '.')
			}
		}
		warehouse = append(warehouse, newLine)
	}

	var moves []byte
	for _, line := range input[movesIdx+1:] {
		moves = append(moves, line...)
	}

	var toMove [][2]int
	visited := map[[2]int]bool{}
	canMove = func(r, c int, isFirst bool) bool {
		if visited[[2]int{r, c}] {
			return true
		}
		visited[[2]int{r, c}] = true
		var ok bool
		switch warehouse[r][c] {
		case '.':
			return true
		case '#':
			return false
		case '@':
			if canMove(r+dr, c+dc, true) {
				toMove = append(toMove, [2]int{r, c})
				return true
			}
		case '[':
			ok = dr == 0 || (isFirst && canMove(r, c+1, false))
		case ']':
			ok = dr == 0 || (isFirst && canMove(r, c-1, false))
		}

		if (ok || !isFirst) && canMove(r+dr, c+dc, true) {
			toMove = append(toMove, [2]int{r, c})
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

		if canMove(row, col, true) {
			row += dr
			col += dc
			for _, m := range toMove {
				r, c := m[0], m[1]
				cell := warehouse[r][c]
				warehouse[r+dr][c+dc] = cell
				warehouse[r][c] = '.'
			}
		}
		toMove = toMove[:0]
		clear(visited)
	}

	for i, row := range warehouse {
		for j, cell := range row {
			if cell == '[' {
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
		// {"../test2.txt", 9021},
		{"../input.txt", 1425169},
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
