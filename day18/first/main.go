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

var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, -1, 1}
var start = Pos{0, 0}

type Pos struct {
	row, col int
}

func parseInput(fileName string, n, m, x int) [][]byte {
	bytes := utils.ReadCSMatrix(fileName)[:x]
	memory := make([][]byte, n)

	for i := range n {
		memory[i] = make([]byte, m)
	}

	for _, b := range bytes {
		r, c := b[1], b[0]
		memory[r][c] = 1
	}

	return memory
}

func rebuildPath(memory [][]byte, prev map[Pos]Pos, start, end Pos) {
	for end != start {
		memory[end.row][end.col] = 2
		end = prev[end]
	}
	memory[end.row][end.col] = 2

	for _, row := range memory {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func neighbors(cur Pos, grid [][]byte) []Pos {
	var res []Pos
	for i := range 4 {
		next := Pos{cur.row + dr[i], cur.col + dc[i]}
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

func solve(memory [][]byte, end Pos) int {
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
			rebuildPath(memory, prev, start, end)
			break
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
	return gScore[end]
}

func main() {
	tests := []struct {
		fileName string
		n, m, x  int
		end      Pos
		want     int
	}{
		{"../test1.txt", 7, 7, 12, Pos{6, 6}, 22},
		{"../input.txt", 71, 71, 1024, Pos{70, 70}, 314},
	}

	for _, test := range tests {
		mem := parseInput(test.fileName, test.n, test.m, test.x)
		got := solve(mem, test.end)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
