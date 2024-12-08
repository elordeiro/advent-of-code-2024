package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	row, col int
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	antennaMap := [][]byte{}
	for scanner.Scan() {
		antennaMap = append(antennaMap, append([]byte{}, scanner.Bytes()...))
	}

	antennaCoords := map[byte][]Coord{}

	for i, row := range antennaMap {
		for j, antenna := range row {
			if antenna == '.' {
				continue
			}
			antennaCoords[antenna] = append(antennaCoords[antenna], Coord{i, j})
		}
	}

	n, m := len(antennaMap), len(antennaMap[0])

	var res int
	visited := map[Coord]bool{}
	for _, coords := range antennaCoords {
		for i, c1 := range coords {
			for _, c2 := range coords[i+1:] {
				dc := c2.col - c1.col
				dr := c2.row - c1.row
				left := Coord{c1.row - dr, c1.col - dc}
				right := Coord{c2.row + dr, c2.col + dc}
				if left.row >= 0 && left.row < n && left.col >= 0 && left.col < m {
					if !visited[left] {
						res++
						visited[left] = true
					}
				}
				if right.row >= 0 && right.row < n && right.col >= 0 && right.col < m {
					if !visited[right] {
						res++
						visited[right] = true
					}
				}
			}
		}
	}

	fmt.Println(res)
}
