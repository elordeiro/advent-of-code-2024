package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
)

type State struct {
	A    int    // Register A
	B    int    // Register B
	C    int    // Register C
	prog []int  // Program Input
	ip   int    // Instruction Pointer
	out  []byte // Program Output
}

var Chrono State

func parseInput(fileName string) {
	input := utils.ReadStringMatrix(fileName)
	Chrono = State{}
	fmt.Sscanf(input[0], "Register A: %d", &Chrono.A)
	fmt.Sscanf(input[1], "Register B: %d", &Chrono.B)
	fmt.Sscanf(input[2], "Register C: %d", &Chrono.C)

	var prog string
	fmt.Sscanf(input[4], "Program: %s", &prog)
	for _, c := range prog {
		if c == ',' {
			continue
		}
		Chrono.prog = append(Chrono.prog, utils.Atoi(string(c)))
	}
}

func comboOp(op int) int {
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

func adv(op int) {
	op = comboOp(op)
	Chrono.A /= int(math.Pow(2, float64(op)))
}

func bdv(op int) {
	op = comboOp(op)
	Chrono.B = Chrono.A / int(math.Pow(2, float64(op)))
}
func cdv(op int) {
	op = comboOp(op)
	Chrono.C = Chrono.A / int(math.Pow(2, float64(op)))
}

func bxl(op int) {
	Chrono.B = Chrono.B ^ op
}

func bst(op int) {
	Chrono.B = comboOp(op) % 8
}

func bxc() {
	Chrono.B = Chrono.B ^ Chrono.C
}

func jnz(op int) {
	if Chrono.A != 0 {
		Chrono.ip = op
	} else {
		Chrono.ip += 2
	}
}

func out(op int) {
	op = comboOp(op) % 8
	Chrono.out = append(Chrono.out, strconv.Itoa(op)...)
	Chrono.out = append(Chrono.out, ',')
}

func solve(fileName string) string {
	parseInput(fileName)

	for Chrono.ip+1 < len(Chrono.prog) {
		switch Chrono.prog[Chrono.ip] {
		case 0:
			adv(Chrono.prog[Chrono.ip+1])
		case 1:
			bxl(Chrono.prog[Chrono.ip+1])
		case 2:
			bst(Chrono.prog[Chrono.ip+1])
		case 3:
			jnz(Chrono.prog[Chrono.ip+1])
			continue
		case 4:
			bxc()
		case 5:
			out(Chrono.prog[Chrono.ip+1])
		case 6:
			bdv(Chrono.prog[Chrono.ip+1])
		case 7:
			cdv(Chrono.prog[Chrono.ip+1])
		}
		Chrono.ip += 2
	}

	return string(Chrono.out[:len(Chrono.out)-1])
}

func main() {
	tests := []struct {
		fileName string
		want     string
	}{
		{"../test1.txt", "4,6,3,5,6,3,5,2,1,0"},
		// {"../test2.txt", "0,1,2"},
		// {"../test3.txt", "4,2,5,6,7,7,7,7,3,1,0"},
		// {"../test4.txt", "0,1,2"},
		// {"../input.txt", ""},
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
