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

	var area, perimeter int
	var plant byte
	n, m := len(garden), len(garden[0])

	visited := map[Cell]bool{}
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			perimeter++
			return
		}
		if garden[i][j] != plant {
			if !visited[Cell{i, j}] {
				perimeter++
			}
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
			area, perimeter = 0, 0
			plant = garden[i][j]
			dfs(i, j)
			res += area * perimeter
			clear(visited)
		}
	}

	fmt.Println(res)
}
