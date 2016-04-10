package main

import (
	"fmt"
	"strings"
)

//T tests
var T int

//N is number seed
var line string
var ints []bool

var stack []bool

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Scan(&line)
		fmt.Printf("Case #%d: %d\n", test, eval())
	}
}

func eval() int {
	lineToStack()
	finalIndex := len(stack) - 1
	flips := 0
	for i := range stack {
		if i == finalIndex {
			break
		}
		if stack[i] != stack[i+1] {
			flips++
		}
	}
	if stack[finalIndex] == false {
		return flips + 1
	}
	return flips
}

func lineToStack() {
	line = strings.TrimSpace(line)
	stack = make([]bool, len(line))
	for i, c := range line {
		stack[i] = c == '+'
	}
}
