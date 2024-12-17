package main

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
	"math"
)

const (
	North = iota
	East
	South
	West
	inf = math.MaxInt
)

type State struct {
	pos   Pos
	score int
	path  string
}

type Pos struct {
	x, y, d int
}

type PathHeap []State

func (h PathHeap) Len() int            { return len(h) }
func (h PathHeap) Less(i, j int) bool  { return h[i].score < h[j].score }
func (h PathHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PathHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *PathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Directions (N, E, S, W)
var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func solve(fileName string) int {
	maze := utils.ReadMatrix(fileName)
	visited := map[Pos]int{}
	paths := []string{}

	hp := &PathHeap{}
	heap.Init(hp)

	n, m := len(maze), len(maze[0])
	start := Pos{n - 2, 1, East}
	end := Pos{x: 1, y: m - 2}

	heap.Push(hp, State{start, 0, ""})
	visited[start] = 0
	maxScore := inf

	for hp.Len() > 0 {
		s := heap.Pop(hp).(State)

		if s.score > maxScore {
			break
		}

		if scr, ok := visited[s.pos]; ok && scr < s.score {
			continue
		}

		visited[s.pos] = s.score
		if s.pos.x == end.x && s.pos.y == end.y {
			maxScore = s.score
			paths = append(paths, s.path)
		}

		dir := s.pos.d
		nx, ny := s.pos.x+dx[dir], s.pos.y+dy[dir]
		if maze[nx][ny] != '#' {
			heap.Push(hp, State{Pos{nx, ny, dir}, s.score + 1, s.path + "F"})
		}

		l, r := (dir+1)%4, (((dir - 1) + 4) % 4)
		heap.Push(hp, State{Pos{s.pos.x, s.pos.y, l}, s.score + 1000, s.path + "R"})
		heap.Push(hp, State{Pos{s.pos.x, s.pos.y, r}, s.score + 1000, s.path + "L"})
	}

	tiles := map[Pos]struct{}{start: {}}
	for _, p := range paths {
		tile := start
		dir := East
		for _, c := range p {
			switch c {
			case 'L':
				dir = ((((dir - 1) % 4) + 4) % 4)
			case 'R':
				dir = ((((dir + 1) % 4) + 4) % 4)
			case 'F':
				tile.x += dx[dir]
				tile.y += dy[dir]
				tiles[tile] = struct{}{}
			}
		}
	}

	return len(tiles)
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 45},
		{"../test2.txt", 64},
		// {"../test3.txt", 21148},
		// {"../test4.txt", 5078},
		// {"../test5.txt", 4013},
		{"../input.txt", 143564},
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
