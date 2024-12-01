package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(num int, slice []int) int {
	total := 0
	for _, e := range slice {
		if e == num {
			total++
		}
	}
	return total
}

func main() {
	llist := make([]int, 1000)
	rlist := make([]int, 1000)
	counts := make([]int, 1000)

	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%d %d", &llist[i], &rlist[i])
		i++
	}

	for i, e := range llist {
		counts[i] = e * count(e, rlist)
	}

	sum := 0
	for _, val := range counts {
		sum += val
	}
	fmt.Println(sum)
}
