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
	Loop
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
	var start Cell
	for scanner.Scan() {
		room = append(room, append([]byte{}, scanner.Bytes()...))
		row := len(room) - 1
		if col := slices.Index(room[row], '^'); col != -1 {
			start = Cell{row, col}
			room[row][col] = Up
		}
	}

	var roomCpy [][]byte
	var cur Cell
	n, m := len(room), len(room[0])
	dir := byte(Up)

	nextStep := func(row, col int) int {
		if row < 0 || col < 0 || row >= n || col >= m {
			return OutOfBounds
		}
		if roomCpy[row][col] == '#' {
			return Blocked
		}

		cur = Cell{row, col}

		if roomCpy[row][col] == '.' {
			roomCpy[row][col] = dir
			return NewCell
		}
		if roomCpy[row][col] == dir {
			return Loop
		}
		return Visited
	}

	var loopCount int
	for row := range n {
		for col := range m {
			if room[row][col] != '.' {
				continue
			}
			roomCpy = clone(room)
			roomCpy[row][col] = '#'
			dir = Up
			cur = start

			for {
				var step int

				switch dir {
				case Up:
					step = nextStep(cur.row-1, cur.col)
					if step == Blocked {
						dir = Right
					}
				case Down:
					step = nextStep(cur.row+1, cur.col)
					if step == Blocked {
						dir = Left
					}
				case Left:
					step = nextStep(cur.row, cur.col-1)
					if step == Blocked {
						dir = Up
					}
				case Right:
					step = nextStep(cur.row, cur.col+1)
					if step == Blocked {
						dir = Down
					}
				}

				if step == OutOfBounds {
					break
				}
				if step == Loop {
					loopCount++
					break
				}
			}
		}
	}

	fmt.Println(loopCount)
}

func clone(room [][]byte) [][]byte {
	cln := make([][]byte, len(room))
	m := len(room[0])
	for i, row := range room {
		cln[i] = make([]byte, m)
		copy(cln[i], row)
	}
	return cln
}
