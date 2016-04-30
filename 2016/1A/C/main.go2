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
		tt := test(i)
		best = max(best, tt)
		//	fmt.Println()
		//	fmt.Println("BEST:\t", best)

	}
	fmt.Print(best)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test(i int) int {
	total := 0
	inspect := i
	alreadyInCircle := make([]bool, N+1)
	history := make([]int, N+1)
	iteration := 0
	//fmt.Println()
	for BFF[inspect] != i && alreadyInCircle[inspect] == false {
		//	fmt.Printf("%d->", inspect)
		history[inspect] = iteration
		iteration++
		alreadyInCircle[inspect] = true
		inspect = BFF[inspect]
		total++
	}
	//	fmt.Println(inspect)
	if BFF[inspect] == i {
		total++
	}
	if inspect == BFF[BFF[inspect]] {
		alreadyInCircle[inspect] = true
		alreadyInCircle[BFF[BFF[inspect]]] = true

		currentBest := total + testAllMini(alreadyInCircle, alreadyInCircle, inspect)
		Best := total + reverseCheck(BFF[inspect], alreadyInCircle)
		currentBest = max(Best, currentBest)
		//fmt.Println(total, currentBest, Best)
		return currentBest
	}
	return total - history[inspect] + testAllMini(alreadyInCircle, alreadyInCircle, -1)
}

func reverseCheck(i int, list []bool) int {
	for j := 1; j <= N; j++ {
		if list[j] == false && BFF[j] == i {
			list[j] = true
			return 1 + reverseCheck(j, list)
		}
	}
	return 0
}

func testAllMini(list, list2 []bool, usable int) int {
	n := 0
	for i := range list {
		list[i] = list[i] || list2[i]
		if list[i] == false {
			n++
		}
	}
	if n == 0 {
		return 0
	}
	best := 0
	for j := 1; j <= n; j++ {
		if list[j] == true {
			continue
		}
		best = max(best, testmini(j, n, list, usable))
	}
	return best
}

func testmini(i, n int, list []bool, usable int) int {
	total := 0
	inspect := i
	alreadyInCircle := make([]bool, N+1)
	history := make([]int, N+1)
	iteration := 0
	//	fmt.Print("\t")
	for BFF[inspect] != i && (alreadyInCircle[inspect] == false || inspect == usable) && (list[inspect] == false || inspect == usable) {
		//	fmt.Printf("%d->", inspect)
		history[inspect] = iteration
		iteration++
		alreadyInCircle[inspect] = true
		inspect = BFF[inspect]
		total++
	}
	if list[inspect] == true {
		return 0
	}
	//fmt.Println(inspect)
	if BFF[inspect] == i {
		total++
	}
	if inspect != BFF[BFF[inspect]] {
		return 0
	}
	if i != BFF[BFF[i]] {
		return 0
	}
	return total //+ testAllMini(alreadyInCircle, list, inspect)

}
