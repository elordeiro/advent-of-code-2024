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

	for i, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			panic(fmt.Sprintf("Failed Test %d\n\tGot %d, Want %d", i+1, got, test.want))
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
