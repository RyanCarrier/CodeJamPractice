package main

import "fmt"

//T test cases
var T int

var line string

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
}
