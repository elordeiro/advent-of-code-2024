package main

import (
	"adventofcode/utils"
	"fmt"
)

func parseInput(fileName string) ([][5]int, [][5]int) {
	input := utils.ReadStringSlice(fileName)

	var locks, keys [][5]int

	for i := 0; i < len(input); {
		line := input[i]
		if line == "" {
			continue
		}

		var curr [5]int
		for j := i + 1; j < i+6; j++ {
			for k, c := range input[j] {
				if c == '#' {
					curr[k]++
				}
			}
		}
		if line[0] == '#' {
			locks = append(locks, curr)
		} else {
			keys = append(keys, curr)
		}
		i += 8
	}

	return locks, keys
}

func solve(fileName string) int {
	locks, keys := parseInput(fileName)

	var res int
	for _, lock := range locks {
	Mid:
		for _, key := range keys {
			for i := range 5 {
				if lock[i]+key[i] > 5 {
					continue Mid
				}
			}
			res++
		}
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 3},
		{"../input.txt", 3107},
	}

	for _, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
