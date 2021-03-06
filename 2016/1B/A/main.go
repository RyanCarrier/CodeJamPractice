package main

import (
	"fmt"
	"strconv"
)

var tests, n int
var result string
var line string

func main() {
	fmt.Scan(&tests)
	for test := 1; test <= tests; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	result = ""
	fmt.Scan(&line)
	list := make([]bool, len(line))
	number(list, 0)
	fmt.Println(result)
}

func number(list []bool, i int) bool {
	//fmt.Println(i, list)
	if alltru(list) || i > 9 {
		return true
	}
	_, ok := removeSwitch(list, i)
	if ok {
		//	fmt.Println(i, result)
		result += strconv.Itoa(i)
		return number(list, i)
	}
	return number(list, i+1)
}

func alltru(list []bool) bool {
	for _, b := range list {
		if b == false {
			return false
		}
	}
	return true
}

func removeSwitch(list []bool, remove int) ([]bool, bool) {
	var s string
	switch remove {
	case 0:
		s = "ZERO"
	case 1:
		s = "ONE"
	case 2:
		s = "TWO"
	case 3:
		s = "THREE"
	case 4:
		s = "FOUR"
	case 5:
		s = "FIVE"
	case 6:
		s = "SIX"
	case 7:
		s = "SEVEN"
	case 8:
		s = "EIGHT"
	case 9:
		s = "NINE"
	}
	return removeALoop(list, s)
}

func removeALoop(list []bool, remove string) ([]bool, bool) {
	var i int
	add := make([]int, len(remove))
	for j, r := range remove {
		i = removeA(list, r)
		if i < 0 {
			return []bool{}, false
		}
		add[j] = i
	}
	for j := range add {
		list[add[j]] = true
	}
	return list, true
}

func removeA(list []bool, r rune) int {
	for i := range list {
		if list[i] == false && rune(line[i]) == r {
			return i
		}
	}
	//	fmt.Println("COULDNT FIND", string(r))
	return -1
}
