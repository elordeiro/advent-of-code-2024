package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const (
	Up = iota
	Down
	Left
	Right

	OutOfBounds
	NewCell
	Visited
	Blocked
)

type Cell struct {
	row, col int
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	room := [][]byte{}
	var curr Cell
	for scanner.Scan() {
		room = append(room, append([]byte{}, scanner.Bytes()...))
		row := len(room) - 1
		if col := slices.Index(room[row], '^'); col != -1 {
			curr = Cell{row, col}
			room[row][col] = 'X'
		}
	}

	visitedCount := 1
	n, m := len(room), len(room[0])
	dir := Up

	nextStep := func(row, col int) int {
		if row < 0 || col < 0 || row >= n || col >= m {
			return OutOfBounds
		}
		if room[row][col] == '#' {
			return Blocked
		}

		curr = Cell{row, col}

		if room[row][col] == 'X' {
			return Visited
		}
		room[row][col] = 'X'
		return NewCell
	}

	for {
		var step int

		switch dir {
		case Up:
			step = nextStep(curr.row-1, curr.col)
			if step == Blocked {
				dir = Right
			}
		case Down:
			step = nextStep(curr.row+1, curr.col)
			if step == Blocked {
				dir = Left
			}
		case Left:
			step = nextStep(curr.row, curr.col-1)
			if step == Blocked {
				dir = Up
			}
		case Right:
			step = nextStep(curr.row, curr.col+1)
			if step == Blocked {
				dir = Down
			}
		}

		if step == OutOfBounds {
			break
		}
		if step == NewCell {
			visitedCount++
		}
	}

	fmt.Println(visitedCount)
}
