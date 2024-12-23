package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

type Coord struct {
	x, y int
}

type Claw struct {
	aOffset, bOffset, prize Coord
}

func solve(fileName string) int {
	input := utils.ReadStringSlice(fileName)

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
		px := utils.Atoi(line[strings.Index(line, "X")+2 : strings.Index(line, ",")])
		py := utils.Atoi(line[strings.Index(line, "Y")+2:])

		claws = append(claws, Claw{Coord{ax, ay}, Coord{bx, by}, Coord{px, py}})
	}

	var res int
	for _, claw := range claws {
		var matches []int
		xmat := make([][]int, 100)
		ymat := make([][]int, 100)
		for i := range 100 {
			xmat[i] = make([]int, 100)
			ymat[i] = make([]int, 100)
		}
		for i := 0; i < 100; i++ {
			xmat[0][i] = claw.aOffset.x * i
			xmat[i][0] = claw.bOffset.x * i
			ymat[0][i] = claw.aOffset.y * i
			ymat[i][0] = claw.bOffset.y * i
		}
		for i := 1; i < 100; i++ {
			x1 := xmat[i][0]
			y1 := ymat[i][0]
			for j := 1; j < 100; j++ {
				x2 := xmat[0][j]
				y2 := ymat[0][j]
				X, Y := x1+x2, y1+y2
				xmat[i][j] = X
				ymat[i][j] = Y

				if X == claw.prize.x && Y == claw.prize.y {
					matches = append(matches, i+j*3)
				}
			}
		}
		slices.Sort(matches)
		if len(matches) > 0 {
			res += matches[0]
		}
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 480},
		{"../input.txt", 30973},
	}

	for i, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			panic(fmt.Sprintf("Failed Test %d\n\tGot %d, Want %d", i+1, got, test.want))
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
