package main

import (
	"bufio"
	"os"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {

	}
}
