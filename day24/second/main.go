package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

const (
	AND = iota
	OR
	XOR
)

type Gate struct {
	in1, in2, out string
	typ           int
}

func parseInput(fileName string) []Gate {
	input := utils.ReadStringSlice(fileName)

	gates := []Gate{}

	var idx int
	for _, line := range input {
		idx++
		if line == "" {
			break
		}
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
		if in2 > in1 {
			in1, in2 = in2, in1
		}

		gates = append(gates, Gate{in1, in2, out, typ})
	}

	return gates
}

func solve(fileName string) string {
	gates := parseInput(fileName)

	var edges []string
	var andID, orID, xorID int
	for _, gate := range gates {
		var gateName string
		var gateID int

		switch gate.typ {
		case AND:
			gateID = andID
			gateName = "AND"
			andID++
		case OR:
			gateID = orID
			orID++
			gateName = "OR"
		case XOR:
			gateID = xorID
			xorID++
			gateName = "XOR"
		}

		edges = append(edges, fmt.Sprintf("%s --> %s%d(\"%s\") --> %s", gate.in1, gateName, gateID, gateName, gate.out))
		edges = append(edges, fmt.Sprintf("%s --> %s%d(\"%s\")", gate.in2, gateName, gateID, gateName))
	}

	slices.SortFunc(edges, func(a, b string) int {
		if a > b {
			return -1
		}
		return 1
	})
	for _, e := range edges {
		fmt.Println(e)
	}

	// Copy output and paste in the live editor for mermaid
	// https://mermaid.live/edit
	// Visualy inspect diagram in order to find swapped wires
	// Swap wires to see if its makes the diagram more 'correct'

	res := []string{"jst", "z05", "gdf", "mcm", "dnt", "z15", "gwc", "z30"}
	slices.Sort(res)
	return strings.Join(res, ",")
}

func main() {
	tests := []struct {
		fileName string
		want     string
	}{
		{"../input.txt", "dnt,gdf,gwc,jst,mcm,z05,z15,z30"},
	}

	for _, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %s, Want %s\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %s\n", test.fileName, got)
	}
}
