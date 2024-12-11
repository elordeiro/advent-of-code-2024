package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  []byte
	next *Node
}

// Linked List method. Works well enough till about 40 "blinks"
func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var head *Node
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	dummy := &Node{}
	curr := dummy
	for _, part := range parts {
		curr.next = &Node{val: []byte(part)}
		curr = curr.next
	}
	head = dummy.next

	var count int
	for range 25 {
		count = 0
		curr := head
		for curr != nil {
			count++
			if string(curr.val) == "0" {
				curr.val = []byte{'1'}
			} else if len(curr.val)%2 == 0 {
				newVal := curr.val[:len(curr.val)/2]
				rightVal, _ := strconv.Atoi(string(curr.val[len(curr.val)/2:]))
				right := &Node{val: []byte(strconv.Itoa(rightVal))}
				curr.val = newVal
				right.next = curr.next
				curr.next = right
				count++
				curr = curr.next
			} else {
				val, _ := strconv.Atoi(string(curr.val))
				val *= 2024
				curr.val = []byte(strconv.Itoa(val))
			}
			curr = curr.next
		}
		curr = head
	}
	fmt.Println(count)
}
