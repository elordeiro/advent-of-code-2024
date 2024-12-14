package main

import (
	"adventofcode/utils"
	"fmt"
)

type Robot struct {
	px, py, vx, vy int
}

var kernel = [][]int{
	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1, 0, 0, 0},
	{0, 0, 1, 1, 1, 1, 1, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func solve(n, m int, fileName string) int {
	list := utils.ReadStringMatrix(fileName)

	var robots []Robot
	for _, e := range list {
		var rob Robot
		fmt.Sscanf(e, "p=%d,%d v=%d,%d", &rob.py, &rob.px, &rob.vy, &rob.vx)
		robots = append(robots, rob)
	}

	floorPlan := make([][]int, n)
	for i := range n {
		floorPlan[i] = make([]int, m)
	}

	for _, r := range robots {
		floorPlan[r.px][r.py] += 1
	}

	var res int
Outer:
	for s := 1; ; s++ {
		for i, r := range robots {
			floorPlan[r.px][r.py] -= 1
			r.px = ((((r.px + r.vx) % n) + n) % n)
			r.py = ((((r.py + r.vy) % m) + m) % m)
			floorPlan[r.px][r.py] += 1
			robots[i] = r
		}

		for i := range n - len(kernel) {
			for j := range m - len(kernel[0]) {
				if isMatch(floorPlan, i, j) {
					res = s
					break Outer
				}
			}
		}
	}
	return res
}

func isMatch(floorPlan [][]int, x, y int) bool {
	for i := range len(kernel) {
		for j := range len(kernel[0]) {
			if kernel[i][j] == 1 && floorPlan[x+i][y+j] == 0 {
				return false
			}
		}
	}

	for i := range len(floorPlan) {
		for j := range len(floorPlan[0]) {
			if floorPlan[i][j] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("^")

			}
		}
		fmt.Printf("\n")
	}
	return true
}

func main() {
	tests := []struct {
		fileName string
		n, m     int
		want     int
	}{
		// {"../test1.txt", 7, 11, 12},
		{"../input.txt", 103, 101, 7584},
	}

	for _, test := range tests {
		got := solve(test.n, test.m, test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
