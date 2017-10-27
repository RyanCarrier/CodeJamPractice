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
	N := ioh.GetInt()
	if iter(ioh.GetIntsN(N)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func iter(A []int) bool {
	if len(A) < 2 {
		return true
	}
	L := len(A)
	var P int
	if L%2 == 0 {
		P = L / 2
	} else {
		P = ((L + 1) / 2)
	}
	P--
	max, min := maxMin(A)
	if A[P] == max || A[P] == min {
		return iter(append(A[:P], A[P+1:]...))
	}
	return false
}

func maxMin(A []int) (int, int) {
	max := -1
	min := 999999999
	for _, v := range A {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
