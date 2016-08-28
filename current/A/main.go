package main

import (
	"fmt"

	ioh "github.com/RyanCarrier/CodeJamPractice/iohelper"
)

//T test cases
var T int

//var io = iohelper
func main() {
	T = ioh.GetInt()
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	L, R := ioh.GetInt(), ioh.GetInt()
	if L == 0 && R == 0 {
		fmt.Println(1)
		return
	}
	if min2(L, R) == 0 {
		fmt.Println(0)
		return
	}
	total := 1
	L--
	R--
	//total += min2(L, R) / 3
	for m := min(L, R, 2); m > 0; m = min(L, R, 2) {
		if total%2 == 1 {
			total++
		}
		total += m

		L -= m
		R -= m
	}
	fmt.Println(total)
}

func min(a, b, c int) int {
	return min2(min2(a, b), c)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
