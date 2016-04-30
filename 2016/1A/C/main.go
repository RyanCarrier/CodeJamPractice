package main

import "fmt"

//T t
var T int

//N how many ints
var N int

//BFF is the id of their bff
var BFF []int

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Scan(&N)
		fmt.Printf("Case #%d: ", test)
		solve()
		fmt.Println()
	}
}

func solve() {
	BFF = make([]int, N+1)
	var bff int
	for i := 0; i < N; i++ {
		fmt.Scan(&bff)
		BFF[i+1] = bff
	}
	best := 0
	for i := 1; best < N && i <= N; i++ {
		best = max(best, test(i, make([]bool, N+1)))
	}
	fmt.Print(best)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseCheck(i int, list []bool) int {
	for j := 1; j <= N; j++ {
		if list[j] == false && BFF[j] == i {
			list[j] = true
			return max(1+reverseCheck(j, list), reverseCheck(i, list))
		}
	}
	return 0
}

func test(i int, used []bool) int {
	history := make([]int, N+1)
	kids := 0
	for used[i] == false {
		history[i] = kids
		kids++
		used[i] = true
		i = BFF[i]
	}
	if i == BFF[i] {
		return kids + reverseCheck(i, used) - history[i]
	}
	return kids - history[i]
}
