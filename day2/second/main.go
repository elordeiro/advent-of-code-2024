package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	safe := 0

Outer:
	for scanner.Scan() {
		report := scanner.Text()
		levelsStr := strings.Split(report, " ")

		levels := []int{}
		for _, level := range levelsStr {
			num, _ := strconv.Atoi(level)
			levels = append(levels, num)
		}

		isDecreasing := false
		if levels[1] < levels[0] {
			isDecreasing = true
		}

		removed := false
		for i := 1; i < len(levels); i++ {
			if isDecreasing {
				if levels[i] < levels[i-1]-3 || levels[i] >= levels[i-1] {
					if removed {
						continue Outer
					}
					removed = true
				}
			} else {
				if levels[i] > levels[i-1]+3 || levels[i] <= levels[i-1] {
					if removed {
						continue Outer
					}
					removed = true
				}
			}
		}
		safe += 1
	}
	fmt.Println(safe)
}
