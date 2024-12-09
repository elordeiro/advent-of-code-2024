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

	diskMap := []byte{}
	scanner.Scan()
	diskMap = append(diskMap, scanner.Bytes()...)

	var fileID int
	var decoded []int
	isFile := true

	for _, fileOrSpace := range diskMap {
		fsLen := fileOrSpace - '0'
		if isFile {
			isFile = false
			for range fsLen {
				decoded = append(decoded, fileID)
			}
			fileID++
		} else {
			isFile = true
			for range fsLen {
				decoded = append(decoded, -1)
			}
		}
	}

	var checkSum int
	l, r := 0, len(decoded)-1
	for l < r {
		for decoded[l] >= 0 {
			checkSum += decoded[l] * l
			l++
		}
		for decoded[r] < 0 {
			r--
		}
		decoded[l] = decoded[r]
		decoded[r] = -1
	}

	fmt.Println(checkSum)
}
