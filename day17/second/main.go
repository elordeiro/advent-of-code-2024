package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"slices"
)

type State struct {
	A    uint64 // Register A
	B    uint64 // Register B
	C    uint64 // Register C
	ip   uint64 // Instruction Pointer
	prog []byte // Program Input
	out  []byte // Program Output
}

var Chrono State
var OrgChrono State

func parseInput(fileName string) {
	input := utils.ReadStringMatrix(fileName)
	OrgChrono = State{}
	fmt.Sscanf(input[0], "Register A: %d", &OrgChrono.A)
	fmt.Sscanf(input[1], "Register B: %d", &OrgChrono.B)
	fmt.Sscanf(input[2], "Register C: %d", &OrgChrono.C)

	var prog string
	fmt.Sscanf(input[4], "Program: %s", &prog)
	for _, c := range prog {
		if c == ',' {
			continue
		}
		OrgChrono.prog = append(OrgChrono.prog, byte(c)-'0')
	}
}

func comboOp(op uint64) uint64 {
	if op > 3 {
		switch op {
		case 4:
			op = Chrono.A
		case 5:
			op = Chrono.B
		case 6:
			op = Chrono.C
		}
	}
	return op
}

func adv(op uint64) {
	op = comboOp(op)
	Chrono.A = Chrono.A / uint64(math.Pow(2, float64(op)))
}

func bdv(op uint64) {
	op = comboOp(op)
	Chrono.B = Chrono.A / uint64(math.Pow(2, float64(op)))
}
func cdv(op uint64) {
	op = comboOp(op)
	Chrono.C = Chrono.A / uint64(math.Pow(2, float64(op)))
}

func bxl(op uint64) {
	Chrono.B = Chrono.B ^ op
}

func bst(op uint64) {
	Chrono.B = comboOp(op) % 8
}

func bxc() {
	Chrono.B = Chrono.B ^ Chrono.C
}

func jnz(op uint64) {
	if Chrono.A != 0 {
		Chrono.ip = op
	} else {
		Chrono.ip += 2
	}
}

func out(op uint64) {
	op = comboOp(op) % 8
	Chrono.out = append(Chrono.out, byte(op))
}

func run(init uint64) string {
	Chrono = OrgChrono
	Chrono.A = init
	for Chrono.ip+1 < uint64(len(Chrono.prog)) {
		opr := uint64(Chrono.prog[Chrono.ip+1])
		switch Chrono.prog[Chrono.ip] {
		case 0:
			adv(opr)
		case 1:
			bxl(opr)
		case 2:
			bst(opr)
		case 3:
			jnz(opr)
			continue
		case 4:
			bxc()
		case 5:
			out(opr)
		case 6:
			bdv(opr)
		case 7:
			cdv(opr)
		}
		Chrono.ip += 2
	}

	return string(Chrono.out)
}

func solve(a uint64, i int) uint64 {
	out := run(a)
	if out == string(Chrono.prog) {
		// fmt.Println(a)
		return a
	}
	if i == 0 || string(out) == string(Chrono.prog)[len(Chrono.prog)-i:] {
		for ni := range 8 {
			if na := solve(8*a+uint64(ni), i+1); na > 0 {
				return na
			}
		}
	}
	return 0
}

func byteArrStr(bytes []byte) string {
	for i := range bytes {
		bytes[i] += '0'
	}
	return string(bytes)
}

func main() {
	tests := []struct {
		fileName string
	}{
		{"../test1.txt"},
		{"../input.txt"},
	}

	for _, test := range tests {
		parseInput(test.fileName)
		want := OrgChrono.prog
		A := solve(0, 0)
		run(A)
		got := Chrono.prog
		if !slices.Equal(got, want) {
			fmt.Printf("Failed Test %s\n\tGot %s, Want %s\n", test.fileName, byteArrStr(got), byteArrStr(want))
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, A)
	}
}
