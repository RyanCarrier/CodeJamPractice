package main

import "fmt"

var tests, n int

var line string

func main() {
	fmt.Scan(&tests)
	for test := 1; test <= tests; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	fmt.Scan(&line)
	list := make([]bool, len(line))
	answer, _ := number(list, 0)
	fmt.Println(answer)
}

func number(list []bool, i int) bool {
	//	fmt.Println("Checking ", i)
	//	fmt.Println(list)
	if alltru(list) {
		return true
	}
	newlist := make([]bool, len(list))
	for i, l := range list {
		newlist[i] = l
	}
	temp, ok := removeSwitch(newlist, i)
	if ok {
		fmt.Print(i)
	}
	if i < 9 {
		return number(list, i+1)
	}
	return "Not ok", false
}

func alltru(list []bool) bool {
	for _, b := range list {
		if b == false {
			return false
		}
	}
	return true
}

func removeALoop(list []bool, remove string) ([]bool, bool) {
	var i int
	newList := list
	for _, r := range remove {
		i = removeA(newList, r)
		if i < 0 {
			//fmt.Println("CONFIRMED", i, string(r))
			return []bool{}, false
		}
		newList[i] = true
	}
	return newList, true
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

func removeA(list []bool, r rune) int {
	for i := range list {
		if list[i] == true {
			continue
		}
		if rune(line[i]) == r {
			return i
		}
	}
	//	fmt.Println("COULDNT FIND", string(r))
	return -1
}
