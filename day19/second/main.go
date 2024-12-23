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
	var dfs func([]byte) int
	visited := map[string]int{}

	dfs = func(partial []byte) int {
		if string(partial) == target {
			return 1
		}
		if count, ok := visited[string(partial)]; ok {
			return count
		}

		var res int
		oldLen := len(partial)
		for _, towel := range towels {
			partial = append(partial, towel...)
			if strings.HasPrefix(target, string(partial)) {
				res += dfs(partial)
			}
			partial = partial[:oldLen]
		}

		visited[string(partial)] = res
		return res
	}

	for _, d := range designs {
		clear(visited)
		target = d
		res += dfs([]byte{})
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 16},
		{"../input.txt", 769668867512623},
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
