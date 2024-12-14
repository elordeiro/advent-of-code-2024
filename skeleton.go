package main

import (
	"fmt"
)

func solve(fileName string) int {
	return 0
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 0},
		{"../test2.txt", 0},
		{"../test3.txt", 0},
		{"../input.txt", 0},
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
