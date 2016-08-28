package main

import (
	"fmt"

	ioh "github.com/RyanCarrier/CodeJamPractice/iohelper"
)

//T test cases
var T int

//Taken is taken
var Taken []bool

//Taken2 fuck u
var Taken2 []bool

//Llist is gay
var Llist []int

//Rlist everything is gay
var Rlist []int

//var io = iohelper
func main() {
	T = ioh.GetInt()
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {

	N, L, R, A, B, c, C, M := ioh.GetInt(), ioh.GetInt(), ioh.GetInt(), ioh.GetInt(), ioh.GetInt(), ioh.GetInt(), ioh.GetInt(), ioh.GetInt()

	Taken = make([]bool, max(max(L, R), M)+1)
	Taken2 = make([]bool, max(max(L, R), M)+1)

	Llist = make([]int, N+1)
	Rlist = make([]int, N+1)

	x, y := L, R
	Llist[1], Rlist[1] = minmax(x, y)
	intervalAdd(minmax(x, y))
	for i := 2; i <= N; i++ {
		x, y = (A*x+B*y+c)%M, (A*y+B*x+C)%M
		Llist[i], Rlist[i] = minmax(x, y)
		//fmt.Print(".")
		intervalAdd(minmax(x, y))
	}
	fmt.Println(total(Llist, N))
}

func intervalAdd(a, b int) {
	for ; a <= b; a++ {
		if !Taken2[a] {
			Taken2[a], Taken[a] = Taken[a], true
		}
	}
}

func total(list []int, N int) int {
	sum := 0
	best := 0
	for i := 1; i <= N; i++ {
		total := 0
		for a, b := minmax(Llist[i], Rlist[i]); a <= b; a++ {
			if Taken[a] && !Taken2[a] {
				total++
			}
		}
		if total > best {
			best = total
		}
		//fmt.Print("-")
	}
	for _, i := range Taken {
		if i {
			sum++
		}
	}
	//take out the most unique one
	return sum - best
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minmax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
