package main

import (
	"fmt"
	"sort"
)

//T t
var T int

//N rows
var N int

var line string

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
		fmt.Println()
	}
}

func solve() {
	fmt.Scan(&N)
	Heights := make([]int, 2501)
	Total := make([]int, 2*N*N)
	var h int
	for i := 0; i < N*(2*N-1); i++ {
		fmt.Scan(&h) //Probs could have done 1 line
		Heights[h]++
		Total[i] = h
	}
	found := 0
	missing := make([]int, N)
	for i := 1; i <= 2500 && found < N; i++ {
		if Heights[i]%2 == 1 {
			missing[found] = i
			found++
		}
	}
	//must be in order, and they won't give us 2 same values
	sort.Ints(missing)
	for i := range missing {
		fmt.Printf("%d ", missing[i])
	}
}
