package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		os.Exit(1)
	}

	indexMap := map[int]byte{
		0: 'm',
		1: 'u',
		2: 'l',
		3: '(',
	}
	doMap := map[int]byte{
		0: 'd',
		1: 'o',
		2: '(',
		3: ')',
	}
	dontMap := map[int]byte{
		0: 'd',
		1: 'o',
		2: 'n',
		3: '\'',
		4: 't',
		5: '(',
		6: ')',
	}

	N := len(input)
	enabled := true
	sum := 0

	for i := 0; i < N; i++ {
		if input[i] == 'd' {
			if enabled {
				j := i
				for j < N && j < i+7 && input[j] == dontMap[j-i] {
					j++
				}
				if j == i+7 {
					enabled = false
				}
				i = j - 1
			} else {
				j := i
				for ; j < N && j < i+4 && input[j] == doMap[j-i]; j++ {
				}
				if j == i+4 {
					enabled = true
				}
				i = j - 1
			}
			continue
		}
		if !enabled {
			continue
		}
		if input[i] == 'm' {
			j := i
			for ; j < N && j < i+4 && input[j] == indexMap[j-i]; j++ {
			}
			if j != i+4 {
				i = j - 1
				continue
			}
			l := j
			for ; j < N && input[j] >= '0' && input[j] <= '9'; j++ {
			}
			if j == N || input[j] != ',' {
				i = j - 1
				continue
			}
			r := j
			firstStr := input[l:r]
			j++
			l = j
			for ; j < N && input[j] >= '0' && input[j] <= '9'; j++ {
			}
			if j == N || input[j] != ')' {
				i = j - 1
				continue
			}
			r = j
			secondStr := input[l:r]
			first, _ := strconv.Atoi(string(firstStr))
			second, _ := strconv.Atoi(string(secondStr))
			sum += first * second
			i = j - 1
		}
	}
	fmt.Println(sum)
}
