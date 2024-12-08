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
	antinodes := [][]byte{}

	var res int

	for i, row := range antennaMap {
		blanks := []byte{}
		for j, antenna := range row {
			if antenna == '.' {
				blanks = append(blanks, '.')
				continue
			}
			blanks = append(blanks, '#')
			antennaCoords[antenna] = append(antennaCoords[antenna], Coord{i, j})
		}
		antinodes = append(antinodes, blanks)
	}

	n, m := len(antennaMap), len(antennaMap[0])

	for _, coords := range antennaCoords {
		for i, c1 := range coords {
			for _, c2 := range coords[i+1:] {
				dc := c2.col - c1.col
				dr := c2.row - c1.row

				left := Coord{c1.row - dr, c1.col - dc}
				for left.row >= 0 && left.row < n && left.col >= 0 && left.col < m {
					antinodes[left.row][left.col] = '#'
					left.row -= dr
					left.col -= dc
				}

				right := Coord{c2.row + dr, c2.col + dc}
				for right.row >= 0 && right.row < n && right.col >= 0 && right.col < m {
					antinodes[right.row][right.col] = '#'
					right.row += dr
					right.col += dc
				}
			}
		}
	}

	for _, row := range antinodes {
		for _, cell := range row {
			if cell == '#' {
				res++
			}
		}
	}

	fmt.Println(res)
}
