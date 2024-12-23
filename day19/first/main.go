package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

func parse(fileName string) ([]string, []string) {
	input := utils.ReadStringSlice(fileName)
	towels := strings.Split(input[0], ",")
	for i, t := range towels {
		towels[i] = strings.TrimSpace(t)
	}
	designs := append([]string{}, input[2:]...)
	return towels, designs
}

func solve(towels, designs []string) int {
	slices.SortFunc(towels, func(a, b string) int {
		return len(b) - len(a)
	})

	var res int
	var target string
	var dfs func([]byte) bool
	visited := map[string]bool{}

	dfs = func(partial []byte) bool {
		if visited[string(partial)] {
			return false
		}
		visited[string(partial)] = true
		if string(partial) == target {
			return true
		}

		oldLen := len(partial)
		for _, towel := range towels {
			partial = append(partial, towel...)
			if strings.HasPrefix(target, string(partial)) {
				if dfs(partial) {
					return true
				}
			}
			partial = partial[:oldLen]
		}
		return false
	}

	for _, d := range designs {
		clear(visited)
		target = d
		if dfs([]byte{}) {
			res++
		}
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 6},
		{"../input.txt", 350},
	}

	for _, test := range tests {
		t, d := parse(test.fileName)
		got := solve(t, d)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
