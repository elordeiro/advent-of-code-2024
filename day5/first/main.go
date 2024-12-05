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

	graph := map[int][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		src, dst := getEdge(text)
		graph[src] = append(graph[src], dst)
	}

	updates := [][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		updates = append(updates, getUpdate(text))
	}

	total := 0

	for _, update := range updates {
		path := map[int]bool{}

	CHECK:
		for _, page := range update {
			for _, neighbor := range graph[page] {
				if _, ok := path[neighbor]; ok {
					break CHECK
				}
			}
			path[page] = true
		}

		if len(path) == len(update) {
			total += update[len(update)/2]
		}
	}

	fmt.Println(total)
}

func getEdge(text string) (int, int) {
	parts := strings.Split(text, "|")
	src, _ := strconv.Atoi(parts[0])
	dst, _ := strconv.Atoi(parts[1])
	return src, dst
}

func getUpdate(text string) []int {
	res := []int{}
	parts := strings.Split(text, ",")
	for _, part := range parts {
		val, _ := strconv.Atoi(part)
		res = append(res, val)
	}
	return res
}
