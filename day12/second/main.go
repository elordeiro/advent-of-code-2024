package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct {
	i, j int
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var garden [][]byte
	for scanner.Scan() {
		garden = append(garden, append([]byte{}, scanner.Bytes()...))
	}

	var area int
	var plant byte
	var dfs func(int, int)
	n, m := len(garden), len(garden[0])
	visited := map[Cell]bool{}

	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || garden[i][j] != plant {
			return
		}
		area++
		garden[i][j] = '#'
		visited[Cell{i, j}] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	var res int
	for i := range n {
		for j := range m {
			if garden[i][j] == '#' {
				continue
			}
			area = 0
			plant = garden[i][j]
			dfs(i, j)
			plot, counted := plot(visited, n, m)
			sides := sides(plot, counted)
			res += area * sides
			clear(visited)
		}
	}

	fmt.Println(res)
}

func plot(visited map[Cell]bool, n, m int) ([][]bool, [][][]bool) {
	n += 2
	m += 2
	grid := make([][]bool, n)
	counted := make([][][]bool, n)
	for i := range n {
		grid[i] = make([]bool, m)
		counted[i] = make([][]bool, m)
		for j := range m {
			counted[i][j] = make([]bool, 4)
		}
	}
	for cell := range visited {
		grid[cell.i+1][cell.j+1] = true
	}
	return grid, counted
}

func sides(grid [][]bool, counted [][][]bool) int {
	n, m := len(grid), len(grid[0])
	var res int

	for i := range n {
		for j := range m {
			if grid[i][j] {
				if !grid[i-1][j] {
					counted[i][j][0] = true
					if !counted[i][j-1][0] {
						res++
					}
				}
				if !grid[i+1][j] {
					counted[i][j][1] = true
					if !counted[i][j-1][1] {
						res++
					}
				}
				if !grid[i][j-1] {
					counted[i][j][2] = true
					if !counted[i-1][j][2] {
						res++
					}
				}
				if !grid[i][j+1] {
					counted[i][j][3] = true
					if !counted[i-1][j][3] {
						res++
					}
				}
			}
		}
	}

	return res
}
