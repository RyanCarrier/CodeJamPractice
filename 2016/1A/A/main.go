package main

import "fmt"

//T t
var T int

//S string
var S string

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
		fmt.Println()
	}
}

func solve() {
	fmt.Scan(&S)
	final := ""
	for i, c := range S {
		//less than [0] go right
		if i == 0 {
			final = string(c)
			continue
		}
		if c < rune(final[0]) {
			final += string(c)
		} else {
			final = string(c) + final
		}
	}
	fmt.Print(final)
}
