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
	defer input.Close()
	scanner := bufio.NewScanner(input)

	testVals := []int{}
	numLists := [][]int{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		tv, _ := strconv.Atoi(parts[0])
		testVals = append(testVals, tv)
		parts = strings.Split(parts[1][1:], " ")
		nums := []int{}
		for _, num := range parts {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		numLists = append(numLists, nums)
	}

	var testVal int
	var numList []int

	var makesTrue func(int, int) bool
	makesTrue = func(acum, i int) bool {
		if acum == testVal && i == len(numList) {
			return true
		}
		if acum > testVal || i >= len(numList) {
			return false
		}

		if makesTrue(acum*numList[i], i+1) {
			return true
		} else if makesTrue(acum+numList[i], i+1) {
			return true
		}
		return false
	}

	var total int
	for i, tv := range testVals {
		testVal = tv
		numList = numLists[i]
		if makesTrue(numList[0], 1) {
			total += tv
		}
	}

	fmt.Println(total)
}
