package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	llist := make([]int, 1000)
	rlist := make([]int, 1000)
	diff := make([]int, 1000)

	input, err := os.Open("../input.txt")
	scanner := bufio.NewScanner(input)
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%d %d", &llist[i], &rlist[i])
		i++
	}

	slices.Sort(llist)
	slices.Sort(rlist)

	for i := range len(llist) {
		diff[i] = int(math.Abs(float64(llist[i] - rlist[i])))
	}

	sum := 0
	for val := range slices.Values(diff) {
		sum += val
	}
	fmt.Println(sum)
}
