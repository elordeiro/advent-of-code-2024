package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadMatrix(fileName string) [][]byte {
	scanner, fp := FileScanner(fileName)
	defer fp.Close()

	var mat [][]byte
	for scanner.Scan() {
		text := scanner.Bytes()
		mat = append(mat, append([]byte{}, text...))
	}
	return mat
}

func ReadStringMatrix(fileName string) []string {
	scanner, fp := FileScanner(fileName)
	defer fp.Close()

	var mat []string
	for scanner.Scan() {
		text := scanner.Text()
		mat = append(mat, text)
	}
	return mat
}

func ReadLine(fileName string) string {
	scanner, fp := FileScanner(fileName)
	defer fp.Close()

	return scanner.Text()
}

func FileScanner(fileName string) (*bufio.Scanner, *os.File) {
	input, err := os.Open(fileName)
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)
	return scanner, input
}

func Atoi(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func AtoiS(slice []string) []int {
	var intSlice []int
	for _, e := range slice {
		intSlice = append(intSlice, Atoi(e))
	}
	return intSlice
}
