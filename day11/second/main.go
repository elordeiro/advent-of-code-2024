package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   string
	level int
}

// Memoization method. Runtime: 42ms
func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	dp := map[Node]int{}
	rocks := strings.Fields(scanner.Text())

	var blink func(string, int) int
	blink = func(rock string, level int) int {
		if res, ok := dp[Node{rock, level}]; ok {
			return res
		}
		if rock == "" {
			return 0
		}
		if level == 0 {
			return 1
		}

		var left string
		var right string

		if rock == "0" {
			left = "1"
		} else if n := len(rock); n%2 == 0 {
			left = rock[:n/2]
			rVal, _ := strconv.Atoi(rock[n/2:])
			right = strconv.Itoa(rVal)
		} else {
			val, _ := strconv.Atoi(rock)
			left = strconv.Itoa(val * 2024)
		}

		dp[Node{rock, level}] = blink(left, level-1) + blink(right, level-1)
		return dp[Node{rock, level}]
	}

	var count int
	for _, rock := range rocks {
		count += blink(rock, 75)
	}

	fmt.Println(count)
}
