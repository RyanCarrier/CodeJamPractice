package main

import "fmt"

//T t
var T int

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d:", test)
		solve()
	}
}

func solve() {

}
