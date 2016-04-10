package main

import (
	"fmt"
	"strconv"
)

//T tests
var T int

//N is number seed
var N int
var ints []bool

var result string

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Scan(&N)
		if N == 0 {
			result = "INSOMNIA"
		} else {
			eval()
		}
		fmt.Printf("Case #%d: %s\n", test, result)
	}
}

func eval() {
	Current := 0
	ints = make([]bool, 10)
	for !all() {
		Current += N
		addNew(Current)
	}
	result = strconv.Itoa(Current)
}

func addNew(i int) {
	s := strconv.Itoa(i)
	for _, c := range s {
		ci, _ := strconv.Atoi(string(c))
		ints[ci] = true
	}
}

func all() bool {
	for _, b := range ints {
		if !b {
			return false
		}
	}
	return true
}
