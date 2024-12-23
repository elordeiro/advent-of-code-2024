package main

import (
	"adventofcode/utils"
	"fmt"
)

type Robot struct {
	px, py, vx, vy int
}

func solve(n, m int, fileName string) int {
	list := utils.ReadStringSlice(fileName)

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

	for _, rob := range robots {
		floorPlan[rob.px][rob.py] += 1
	}

	for range 100 {
		for i, rob := range robots {
			floorPlan[rob.px][rob.py] -= 1
			rob.px = ((((rob.px + rob.vx) % n) + n) % n)
			rob.py = ((((rob.py + rob.vy) % m) + m) % m)
			floorPlan[rob.px][rob.py] += 1
			robots[i] = rob
		}
	}

	var quad1, quad2, quad3, quad4 int
	for i := 0; i < n/2; i++ {
		for j := 0; j < m/2; j++ {
			quad1 += floorPlan[i][j]
			quad2 += floorPlan[i][j+1+(m/2)]
			quad3 += floorPlan[i+1+(n/2)][j]
			quad4 += floorPlan[i+1+(n/2)][j+1+(m/2)]
		}
	}

	return quad1 * quad2 * quad3 * quad4
}

func main() {
	tests := []struct {
		fileName string
		n, m     int
		want     int
	}{
		{"../test1.txt", 7, 11, 12},
		{"../input.txt", 103, 101, 236628054},
	}

	for _, test := range tests {
		got := solve(test.n, test.m, test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d", test.fileName, got, test.want)
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
