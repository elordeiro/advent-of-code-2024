package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
)

const (
	Track   = '.'
	Start   = 'S'
	Wall    = '#'
	End     = 'E'
	Time    = 1
	Visited = 0
)

// Globals -----------------------

// Directions   U, D, L, R
var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, -1, 1}
var start, end Pos
var cutOffTime int

// -------------------------------

type Pos struct {
	row, col int
}

func getFastestTime(raceTrack [][]byte) int {
	var q []Pos
	var p Pos

	m := len(raceTrack[0])
	visited := make([][][2]int, len(raceTrack))
	for i := range len(raceTrack) {
		visited[i] = make([][2]int, m)
	}

	q = append(q, start)
	for len(q) > 0 {
		n := len(q)
		for range n {
			p, q = q[0], q[1:]
			if p == end {
				return visited[end.row][end.col][Time]
			}
			if visited[p.row][p.col][Time] > cutOffTime {
				continue
			}
			visited[p.row][p.col][Visited] = 1
			for i := range 4 {
				next := Pos{p.row + dr[i], p.col + dc[i]}
				if raceTrack[next.row][next.col] == Wall || visited[next.row][next.col][Visited] == 1 {
					continue
				}
				q = append(q, next)
				visited[next.row][next.col][Time] = visited[p.row][p.col][Time] + 1
			}
		}
	}

	return -1
}

func solve(fileName string, diff int) int {
	raceTrack := utils.ReadMatrix(fileName)
	var walls []Pos
	n, m := len(raceTrack), len(raceTrack[0])

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			cell := raceTrack[i][j]
			switch cell {
			case Wall:
				walls = append(walls, Pos{i, j})
			case Start:
				start = Pos{i, j}
			case End:
				end = Pos{i, j}
			}
		}
	}

	var res int
	cutOffTime = math.MaxInt
	cutOffTime = getFastestTime(raceTrack) - diff

	for _, wall := range walls {
		raceTrack[wall.row][wall.col] = Track
		if t := getFastestTime(raceTrack); t != -1 && t <= cutOffTime {
			res++
		}
		raceTrack[wall.row][wall.col] = Wall
	}
	return res
}

func main() {
	tests := []struct {
		fileName string
		diff     int
		want     int
	}{
		{"../test1.txt", 1, 44},
		{"../input.txt", 100, 1327},
	}

	for _, test := range tests {
		got := solve(test.fileName, test.diff)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
