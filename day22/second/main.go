package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
)

const (
	Prune = 16777216
)

type Seq struct {
	a, b, c, d int
}

func solve(fileName string, N int) int {
	buyers := utils.ReadIntSlice(fileName)

	globalSeqs := map[Seq]int{}
	for _, s := range buyers {
		localSeqs := map[Seq]int{}
		var a, b, c, d, e int
		a = s % 10

		for n := range N {
			s = (s<<6 ^ s) % Prune
			s = (s>>5 ^ s) % Prune
			s = (s<<11 ^ s) % Prune

			switch n {
			case 0:
				b = s % 10
			case 1:
				c = s % 10
			case 2:
				d = s % 10
			default:
				e = s % 10
				key := Seq{b - a, c - b, d - c, e - d}
				if _, ok := localSeqs[key]; !ok {
					localSeqs[key] = e
				}
				a, b, c, d = b, c, d, e
			}
		}

		for k, v := range localSeqs {
			globalSeqs[k] += v
		}
	}

	res := math.MinInt
	for _, v := range globalSeqs {
		res = max(res, v)
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		N        int
		want     int
	}{
		{"../test2.txt", 10, 6},
		{"../test3.txt", 2000, 23},
		{"../test4.txt", 2000, 27},
		{"../input.txt", 2000, 2223},
	}

	for _, test := range tests {
		got := solve(test.fileName, test.N)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
