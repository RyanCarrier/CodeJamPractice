package main

import (
	"fmt"
	"math"
)

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
	var M int
	fmt.Scan(&M)
	var c int
	fmt.Scan(&c)
	C := make([]int, M+1)
	C[0] = -c
	for i := 0; i < M; i++ {
		fmt.Scan(&C[i+1])
	}
	l := len(C)
	max := float64(1)
	min := float64(-1)
	ceil := float64(0.99999999999)
	floor := float64(-0.99999999999)
	if tried := intry(C, l, 1); int(tried*10000000000) == 0 {
		fmt.Println(1)
		return
	}
	if tried := intry(C, l, -1); int(tried*10000000000) == 0 {
		fmt.Println(-1)
		return
	}

	for tried := try(C, l, avg(max, min)); int(tried*10000000000) != 0; tried = try(C, l, avg(max, min)) {
		a := avg(max, min)
		if a == max || a == min {
			break
		}
		if tried < 0 {
			max = a
		} else {
			min = a
		}

		if a > ceil {
			fmt.Println(ceil)
			return
		}
		if a < floor {
			fmt.Println(floor)
			return
		}
	}
	fmt.Println(avg(max, min))
	//fmt.Println(C)
}

func try(C []int, l int, r float64) float64 {
	total := float64(0)
	j := 0
	for i := len(C) - 1; i >= 0; i-- {
		total += float64(C[i]) * (math.Pow(1+r, float64(j)))
		j++
	}
	return total
}

func intry(C []int, l int, r int) float64 {
	if r == -1 {
		return float64(C[l-1])
	}
	total := float64(0)
	j := 0
	for i := len(C) - 1; i >= 0; i-- {
		total += float64(C[i]) * (math.Pow(2, float64(j)))
		j++
	}
	return total
}

func avg(a, b float64) float64 {
	return (a + b) / 2
}
