package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	tmap := [][]byte{}
	for scanner.Scan() {
		tmap = append(tmap, append([]byte{}, scanner.Bytes()...))
	}

	n, m := len(tmap), len(tmap[0])

	var score int
	var hike func(int, int, byte)

	hike = func(i, j int, target byte) {
		if i < 0 || i >= n || j < 0 || j >= m || tmap[i][j] != target {
			return
		}
		if target == '9' {
			score++
			return
		}
		target++
		hike(i-1, j, target)
		hike(i+1, j, target)
		hike(i, j-1, target)
		hike(i, j+1, target)
	}

	for i := range n {
		for j := range m {
			if tmap[i][j] != '0' {
				continue
			}
			hike(i, j, '0')
		}
	}

	fmt.Println(score)
}
