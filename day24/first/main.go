package main

import (
	"adventofcode/utils"
	"fmt"
	"maps"
	"slices"
	"strings"
)

const (
	AND = iota
	OR
	XOR
)

type Gate struct {
	out string
	typ int
}

func parseInput(fileName string) (map[string]bool, map[string]map[string][]Gate) {
	input := utils.ReadStringSlice(fileName)

	ins := map[string]bool{}
	gates := map[string]map[string][]Gate{}

	var idx int
	for _, line := range input {
		idx++
		if line == "" {
			break
		}
		parts := strings.Fields(line)
		name := strings.Trim(parts[0], ":")
		v := utils.Atoi(strings.TrimSpace(parts[1]))
		var val bool
		if v == 1 {
			val = true
		}
		ins[name] = val
	}

	for _, line := range input[idx:] {
		parts := strings.Fields(line)

		var typ int
		switch parts[1] {
		case "AND":
			typ = 0
		case "OR":
			typ = 1
		case "XOR":
			typ = 2
		}

		in1, in2, out := parts[0], parts[2], parts[4]
		if gates[in1] == nil {
			gates[in1] = make(map[string][]Gate)
		}
		if gates[in2] == nil {
			gates[in2] = make(map[string][]Gate)
		}

		slice := gates[in1][in2]
		slice = append(slice, Gate{out, typ})
		gates[in1][in2] = slice
		slice = gates[in2][in1]
		slice = append(slice, Gate{out, typ})
		gates[in2][in1] = slice
	}

	return ins, gates
}

func solve(fileName string) int {
	wires, gates := parseInput(fileName)

	var q []string
	for w := range wires {
		q = append(q, w)
	}

	out := map[string]bool{}
	for len(q) > 0 {
		n := len(q)
		for range n {
			in1 := q[0]
			q = q[1:]

			for in2, outs := range gates[in1] {
				for _, gate := range outs {
					var o bool
					switch gate.typ {
					case AND:
						o = wires[in1] && wires[in2]
					case OR:
						o = wires[in1] || wires[in2]
					case XOR:
						o = wires[in1] != wires[in2]
					}
					wires[gate.out] = o
					q = append(q, gate.out)
					if strings.HasPrefix(gate.out, "z") {
						out[gate.out] = o
					}
				}
			}
		}
	}

	bits := slices.Collect(maps.Keys(out))
	slices.SortFunc(bits, func(a, b string) int {
		if b < a {
			return -1
		}
		return 1
	})

	var res int
	for _, b := range bits {
		switch wires[b] {
		case true:
			res = res<<1 | 1
		case false:
			res = res << 1
		}
	}
	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 2024},
		{"../input.txt", 47666458872582},
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
