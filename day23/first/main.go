package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

func newGraph(fileName string) map[string][]string {
	conns := utils.ReadStringSlice(fileName)
	graph := map[string][]string{}
	for _, conn := range conns {
		parts := strings.Split(conn, "-")
		src, dst := parts[0], parts[1]
		sns, dns := graph[src], graph[dst]
		sns = append(sns, dst)
		dns = append(dns, src)
		graph[src], graph[dst] = sns, dns
	}
	return graph
}

func solve(fileName string) int {
	graph := newGraph(fileName)

	cycles := map[string]struct{}{}
	var target, n1, n2 string

	var findCycle func(string, int)
	findCycle = func(curr string, depth int) {
		if depth == 3 {
			if curr == target && (curr[0] == 't' || n1[0] == 't' || n2[0] == 't') {
				cycle := []string{curr, n1, n2}
				slices.Sort(cycle)
				cycles[cycle[0]+cycle[1]+cycle[2]] = struct{}{}
			}
			return
		}
		n2 = curr
		for _, neighbor := range graph[curr] {
			findCycle(neighbor, depth+1)
		}
	}

	for curr, neighbors := range graph {
		target = curr
		for _, n := range neighbors {
			n1 = n
			findCycle(n1, 1)
		}
	}

	return len(cycles)
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 7},
		{"../input.txt", 1314},
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
