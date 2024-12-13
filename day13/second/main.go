package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type Coord struct {
	x, y int
}

type Claw struct {
	a, b, p Coord
}

func solve(fileName string) int {
	input := utils.ReadStringMatrix(fileName)

	var claws []Claw

	for i := 0; i < len(input); i += 4 {
		if input[i] == "" {
			continue
		}
		line := input[i]
		ax := utils.Atoi(line[strings.Index(line, "X")+1 : strings.Index(line, ",")])
		ay := utils.Atoi(line[strings.Index(line, "Y")+1:])

		line = input[i+1]
		bx := utils.Atoi(line[strings.Index(line, "X")+1 : strings.Index(line, ",")])
		by := utils.Atoi(line[strings.Index(line, "Y")+1:])

		line = input[i+2]
		px := utils.Atoi(line[strings.Index(line, "X")+2:strings.Index(line, ",")]) + 10000000000000
		py := utils.Atoi(line[strings.Index(line, "Y")+2:]) + 10000000000000

		claws = append(claws, Claw{Coord{ax, ay}, Coord{bx, by}, Coord{px, py}})
	}

	var res int
	for _, c := range claws {
		b := (c.p.y*c.a.x - c.p.x*c.a.y) / (c.b.y*c.a.x - c.b.x*c.a.y)
		a := (c.p.x - b*c.b.x) / c.a.x

		if a*c.a.x+c.b.x*b != c.p.x || a*c.a.y+c.b.y*b != c.p.y {
			continue
		}

		res += (3 * a) + b
	}
	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		// {"../test1.txt", 480},
		{"../input.txt", 95688837203288},
	}

	for i, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			panic(fmt.Sprintf("Failed Test %d\n\tGot %d, Want %d", i+1, got, test.want))
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
