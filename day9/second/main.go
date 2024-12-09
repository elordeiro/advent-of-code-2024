package main

import (
	"bufio"
	"fmt"
	"os"
)

type Info struct {
	i, len int
}

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

	files := []Info{}
	spaces := []Info{}

	for _, fileOrSpace := range diskMap {
		fsLen := fileOrSpace - '0'
		if isFile {
			isFile = false
			files = append(files, Info{len(decoded), int(fsLen)})
			for range fsLen {
				decoded = append(decoded, fileID)
			}
			fileID++
		} else {
			isFile = true
			if fsLen > 0 {
				spaces = append(spaces, Info{len(decoded), int(fsLen)})
			}
			for range fsLen {
				decoded = append(decoded, -1)
			}
		}
	}

	filesLen := len(files) - 1
	spacesLen := len(spaces)
	for fi := filesLen; fi >= 0; fi-- {
		f := files[fi]
		x := decoded[f.i]
		for si := 0; si < spacesLen; si++ {
			s := spaces[si]
			if s.len >= f.len && s.i < f.i {
				for range f.len {
					decoded[s.i] = x
					decoded[f.i] = -1
					s.i++
					f.i++
				}
				s.len -= f.len
				spaces[si] = s
				break
			}
		}
	}

	var checkSum int
	for i, e := range decoded {
		if e > 0 {
			checkSum += i * e
		}
	}

	fmt.Println(checkSum)
}
