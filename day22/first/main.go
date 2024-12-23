package main

import (
	"adventofcode/utils"
	"fmt"
)

const (
	Prune = 16777216
)

func solve(fileName string) int {
	buyers := utils.ReadIntSlice(fileName)

	var res int

	for i, b := range buyers {
		for range 2000 {
			b = (b<<6 ^ b) % Prune
			b = (b>>5 ^ b) % Prune
			b = (b<<11 ^ b) % Prune
		}
		buyers[i] = b
		res += b
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 37327623},
		{"../test4.txt", 8876699},
		{"../input.txt", 19854248602},
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
