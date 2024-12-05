package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	cols := [][]byte{}
	for c := range m {
		col := make([]byte, n)
		for r := range n {
			col[r] = mat[r][c]
		}
		cols = append(cols, col)
	}

	posdiags := [][]byte{}
	for r := range n {
		diag := []byte{}
		for c := 0; r > -1 && c < m; r, c = r-1, c+1 {
			diag = append(diag, mat[r][c])
		}
		posdiags = append(posdiags, diag)
	}
	for c := m - 1; c > 0; c-- {
		diag := []byte{}
		oldC := c
		for r := n - 1; r > -1 && c < m; r, c = r-1, c+1 {
			diag = append(diag, mat[r][c])
		}
		c = oldC
		posdiags = append(posdiags, diag)
	}

	negdiags := [][]byte{}
	for c := m - 1; c > 0; c-- {
		diag := []byte{}
		oldC := c
		for r := 0; r < n && c < m; r, c = r+1, c+1 {
			diag = append(diag, mat[r][c])
		}
		c = oldC
		negdiags = append(negdiags, diag)
	}
	for c := range m {
		diag := []byte{}
		for r := n - 1; r > -1 && c > -1; r, c = r-1, c-1 {
			diag = append(diag, mat[r][c])
		}
		negdiags = append(negdiags, diag)
	}

	total := 0
	re := regexp.MustCompile("XMAS")
	rb := regexp.MustCompile("SAMX")

	for _, row := range mat {
		total += len(re.FindAll(row, -1))
		total += len(rb.FindAll(row, -1))
	}
	for _, col := range cols {
		total += len(re.FindAll(col, -1))
		total += len(rb.FindAll(col, -1))
	}
	for _, diag := range posdiags {
		total += len(re.FindAll(diag, -1))
		total += len(rb.FindAll(diag, -1))
	}
	for _, diag := range negdiags {
		total += len(re.FindAll(diag, -1))
		total += len(rb.FindAll(diag, -1))
	}

	fmt.Println(total)
}
