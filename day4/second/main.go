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

	mat := [][]byte{}
	for scanner.Scan() {
		mat = append(mat, append([]byte{}, scanner.Bytes()...))
	}
	n, m := len(mat), len(mat[0])

	total := 0
	for r := 1; r < n-1; r++ {
		for c := 1; c < m-1; c++ {
			if mat[r][c] != 'A' {
				continue
			}
			str := string([]byte{mat[r-1][c-1], 'A', mat[r+1][c+1]})
			if str != "MAS" && str != "SAM" {
				continue
			}
			str = string([]byte{mat[r+1][c-1], 'A', mat[r-1][c+1]})
			if str != "MAS" && str != "SAM" {
				continue
			}
			total++
		}
	}

	fmt.Println(total)
}
