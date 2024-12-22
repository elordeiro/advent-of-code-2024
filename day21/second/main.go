package main

import (
	"adventofcode/utils"
	"fmt"
)

const (
	Horizontal = iota
	Vertical
)

type Pos struct {
	row, col int
}

type Memo struct {
	seq   string
	depth byte
}

var dp = map[Memo]int{}  // Memoization map
var nkpStart = Pos{3, 2} // Numerical keypad start Pos
var dkpStart = Pos{0, 2} // Directional keypad start Pos

// Char to Pos on numeric keypad
func ctopn(c byte) Pos {
	if c == '0' {
		return Pos{3, 1}
	}
	if c == 'A' {
		return nkpStart
	}
	row := 2 - ((c - '0' - 1) / 3)
	col := (c - '0' - 1) % 3
	return Pos{int(row), int(col)}
}

// Char to Pos on directional keypad
func ctopd(d byte) Pos {
	switch d {
	case '^':
		return Pos{0, 1}
	case '<':
		return Pos{1, 0}
	case 'v':
		return Pos{1, 1}
	case '>':
		return Pos{1, 2}
	default:
		return Pos{0, 2}
	}
}

func pathWriter(off, dir int) []byte {
	var path []byte
	var c byte
	if dir == Horizontal {
		if off < 0 {
			c = '>'
		} else {
			c = '<'
		}
	} else {
		if off < 0 {
			c = 'v'
		} else {
			c = '^'
		}
	}
	for range utils.Abs(off) {
		path = append(path, c)
	}
	return path
}

func shortestSeq(src, dst Pos, isNumPad bool) []byte {
	var path []byte

	dr := src.row - dst.row
	dc := src.col - dst.col

	movesV := pathWriter(dr, Vertical)
	movesH := pathWriter(dc, Horizontal)

	var onGap bool
	if isNumPad {
		onGap = (src.row == 3 && dst.col == 0) || (src.col == 0 && dst.row == 3)
	} else {
		onGap = (src.col == 0 && dst.row == 0) || (src.row == 0 && dst.col == 0)
	}

	goingLeft := dst.col < src.col

	if goingLeft != onGap {
		movesV, movesH = movesH, movesV
	}

	path = append(append([]byte{}, movesV...), movesH...)
	path = append(path, 'A')
	return path
}

func dfs(memo Memo) int {
	if v, ok := dp[memo]; ok {
		return v
	}
	if memo.depth == 0 {
		return len(memo.seq)
	}

	var res int
	var path [][]byte
	prev := dkpStart
	for _, c := range memo.seq {
		curr := ctopd(byte(c))
		path = append(path, shortestSeq(prev, curr, false))
		prev = curr
	}

	for _, p := range path {
		res += dfs(Memo{string(p), memo.depth - 1})
	}
	dp[memo] = res
	return res
}

func solve(fileName string) int {
	var res int
	codes := utils.ReadMatrix(fileName)

	var codeInt int
	for _, code := range codes {
		var path [][]byte
		codeInt = utils.Atoi(string(code[:len(code)-1]))

		prev := nkpStart
		for _, c := range code {
			curr := ctopn(c)
			path = append(path, shortestSeq(prev, curr, true))
			prev = curr
		}

		var pathLen int
		for _, code := range path {
			pathLen += dfs(Memo{string(code), 25})
		}
		res += pathLen * codeInt
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		want     int
	}{
		{"../test1.txt", 154115708116294},
		{"../input.txt", 228800606998554},
	}

	for _, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
