package main

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
	"math"
)

const (
	North = iota
	West
	South
	East
	inf = math.MaxInt
)

var dirs = []Cell{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

type Heap []Node

func (ph Heap) Len() int      { return len(ph) }
func (ph Heap) Swap(i, j int) { ph[i], ph[j] = ph[j], ph[i] }
func (ph Heap) Less(i, j int) bool {
	return ph[i].score < ph[j].score
}
func (ph *Heap) Push(val any) { *ph = append(*ph, val.(Node)) }
func (ph *Heap) Pop() any {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[0 : n-1]
	return x
}

type Cell struct {
	row, col int
}

type Node struct {
	cell       Cell
	score, dir int
}

func score(to int, from [4]int) int {
	if from[to] != inf {
		return 1
	}
	return 1001
}

func solve(fileName string) int {
	maze := utils.ReadMatrix(fileName)
	n, m := len(maze), len(maze[0])
	start, end := Cell{n - 2, 1}, Cell{1, m - 2}

	distances := map[Cell][4]int{}
	for i := range n {
		for j := range m {
			if maze[i][j] == '#' {
				continue
			}
			distances[Cell{i, j}] = [4]int{inf, inf, inf, inf}
		}
	}
	distances[start] = [4]int{inf, inf, inf, 0}

	hp := Heap{Node{start, 0, East}}
	for len(hp) > 0 {
		top := heap.Pop(&hp).(Node)
		source := top.cell
		if source == end {
			break
		}

		for to, dir := range dirs {
			dest := Cell{source.row + dir.row, source.col + dir.col}
			if maze[dest.row][dest.col] == '#' {
				continue
			}
			dists := distances[dest]
			score := score(to, distances[source]) + top.score
			if dists[to] > score {
				dists[to] = score
				distances[dest] = dists
				heap.Push(&hp, Node{dest, score, to})
			}
		}
	}

	ds := distances[end]
	return min(ds[0], min(ds[1], min(ds[2], ds[3])))
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 7036},
		{"../test2.txt", 11048},
		{"../test3.txt", 21148},
		{"../test4.txt", 5078},
		{"../test5.txt", 4013},
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
