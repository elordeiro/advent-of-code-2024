package main

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
	"math"
)

const (
	inf = math.MaxInt
)

// ----------------------------------------------------------------------------
type Cell struct {
	pos    Pos
	fScore int
}

type Heap []Cell

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i].fScore < h[j].fScore }
func (h *Heap) Push(val any)      { *h = append(*h, val.(Cell)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ----------------------------------------------------------------------------

var dirs = []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var start = Pos{0, 0}

type Pos struct {
	row, col int
}

func parseInput(fileName string, n, m, x int) ([][]byte, [][]int) {
	bytes := utils.ReadCSMatrix(fileName)
	memory := make([][]byte, n)

	for i := range n {
		memory[i] = make([]byte, m)
	}

	for _, b := range bytes[:x] {
		r, c := b[1], b[0]
		memory[r][c] = 1
	}

	return memory, bytes
}

func neighbors(cur Pos, grid [][]byte) []Pos {
	var res []Pos
	for _, dir := range dirs {
		next := Pos{cur.row + dir.row, cur.col + dir.col}
		if next.row < 0 || next.row >= len(grid) || next.col < 0 || next.col >= len(grid[0]) {
			continue
		}
		if grid[next.row][next.col] == 1 {
			continue
		}
		res = append(res, next)
	}
	return res
}

func hasPath(memory [][]byte, end Pos) bool {
	n, m := len(memory), len(memory[0])

	visited := map[Pos]bool{start: true}
	gScore := map[Pos]int{}
	fScore := map[Pos]int{}
	prev := map[Pos]Pos{}

	hp := &Heap{Cell{start, 0}}
	heap.Init(hp)

	f := func(pos Pos) int {
		dx := utils.Abs(pos.row - end.row)
		dy := utils.Abs(pos.col - end.col)
		return gScore[pos] + dx + dy
	}

	for i := range n {
		for j := range m {
			if memory[i][j] != 1 {
				gScore[Pos{i, j}] = inf
				fScore[Pos{i, j}] = inf
			}
		}
	}
	gScore[start] = 0
	fScore[start] = f(start)

	for hp.Len() > 0 {
		cur := heap.Pop(hp).(Cell).pos
		if cur == end {
			return true
		}
		visited[cur] = true
		for _, next := range neighbors(cur, memory) {
			tentative := gScore[cur] + 1
			if tentative < gScore[next] {
				prev[next] = cur
				gScore[next] = tentative
				fScore[next] = f(next)
				if !visited[next] {
					heap.Push(hp, Cell{next, fScore[next]})
				}
			}
		}
	}
	return false
}

func solve(memory [][]byte, bytes [][]int, x int, end Pos) Pos {
	for _, byt := range bytes[x:] {
		pos := Pos{byt[1], byt[0]}
		memory[pos.row][pos.col] = 1
		if !hasPath(memory, end) {
			return pos
		}
	}
	return end
}

func main() {
	tests := []struct {
		fileName string
		n, m, x  int
		end      Pos
		want     Pos
	}{
		{"../test1.txt", 7, 7, 12, Pos{6, 6}, Pos{1, 6}},
		{"../input.txt", 71, 71, 1024, Pos{70, 70}, Pos{20, 15}},
	}

	for _, test := range tests {
		mem, bytes := parseInput(test.fileName, test.n, test.m, test.x)
		got := solve(mem, bytes, test.x, test.end)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %v, Want %v\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %v\n", test.fileName, got)
	}
}
