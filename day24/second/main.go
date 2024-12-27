package main

import (
	"adventofcode/utils"
	"fmt"
	"maps"
	"slices"
	"strings"
)

const (
	Nil = iota
	AND
	OR
	XOR
)

type Gate struct {
	in1, in2 string
	typ      int
}

type Adder struct {
	circuit map[string]Gate
}

func parseInput(fileName string) Adder {
	input := utils.ReadStringSlice(fileName)
	circuit := map[string]Gate{}

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
			typ = AND
		case "OR":
			typ = OR
		case "XOR":
			typ = XOR
		}

		in1, in2, out := parts[0], parts[2], parts[4]
		if in1 > in2 {
			in1, in2 = in2, in1
		}
		circuit[out] = Gate{in1, in2, typ}
	}

	return Adder{circuit}
}

func (a Adder) prevGate(x string) int {
	return a.circuit[x].typ
}

func (a Adder) prevX(x string) string {
	return a.circuit[x].in1
}

// Safe to assume that rule1 is only called for z's
func (a Adder) rule1(out string, gate Gate) bool {
	// All z's come from XOR gates, except z45
	if out == "z45" {
		if gate.typ != OR {
			return false
		}
	} else if gate.typ != XOR {
		return false
	}

	// z00 comes from XOR gate with x00 and y00
	if out == "z00" {
		if gate.typ != XOR || (gate.in1 != "x00" || gate.in2 != "y00") {
			return false
		}
		return true
	}

	// No other z's come 'directly' after x's and y's
	return !strings.HasPrefix(gate.in1, "x") && !strings.HasPrefix(gate.in2, "y")
}

func (a Adder) rule2(out string, gate Gate) string {
	switch gate.typ {
	case OR:
		// OR gates's inputs are always AND gates outputs
		if a.prevGate(gate.in1) != AND {
			return gate.in1
		}
		if a.prevGate(gate.in2) != AND {
			return gate.in2
		}

	case AND:
		// Except for x00, there's never 2 AND gates in a row
		if a.prevGate(gate.in1) == AND && a.prevX(gate.in1) != "x00" {
			return gate.in1
		}
		if a.prevGate(gate.in2) == AND && a.prevX(gate.in2) != "x00" {
			return gate.in2
		}

	case XOR:
		// XOR gates that come from OR and XOR gates always output to z's
		prev1 := a.prevGate(gate.in1)
		prev2 := a.prevGate(gate.in2)
		if prev1 == XOR && prev2 == OR || prev2 == XOR && prev1 == OR {
			if !strings.HasPrefix(out, "z") {
				return out
			}
		}
	}

	return ""
}

func solve(fileName string) string {
	res := map[string]struct{}{}

	adder := parseInput(fileName)
	var q []string
	for i := range 46 {
		q = append(q, fmt.Sprintf("z%02d", i))
	}

	onZs := true
	for len(q) > 0 {
		n := len(q)
		for range n {
			out := q[0]
			q = q[1:]
			if gate, ok := adder.circuit[out]; ok {
				if onZs && !adder.rule1(out, gate) {
					res[out] = struct{}{}
				} else {
					if r := adder.rule2(out, gate); len(r) > 0 {
						res[r] = struct{}{}
					}
				}
				q = append(q, gate.in1, gate.in2)
			}
		}
		onZs = false
	}

	r := slices.Collect(maps.Keys(res))
	slices.Sort(r)
	return strings.Join(r, ",")
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
			fmt.Printf("Failed Test %s\n\tGot %s\n\tWant %s\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %s\n", test.fileName, got)
	}
}
