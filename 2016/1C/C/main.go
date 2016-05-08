package main

import "fmt"

var tests, n int

var j, p, s, k int
var line string
var days [][]int

func main() {
	fmt.Scan(&tests)
	for test := 1; test <= tests; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	days = [][]int{}
	fmt.Scan(&j, &p, &s, &k)
	js := make([][]int, j+1)
	for i := range js {
		js[i] = make([]int, s+1)
	}
	jp := make([][]int, j+1)
	for i := range jp {
		jp[i] = make([]int, p+1)
	}
	ps := make([][]int, p+1)
	for i := range ps {
		ps[i] = make([]int, s+1)
	}

	//O(j*p*s) but fuck you
	for jj := 1; jj <= j; jj++ {
		for pp := 1; pp <= p; pp++ {
			for ss := 1; ss <= s; ss++ {
				if js[jj][ss] < k && ps[pp][ss] < k && jp[jj][pp] < k {
					days = append(days, []int{jj, pp, ss})
					js[jj][ss]++
					jp[jj][pp]++
					ps[pp][ss]++
				}
			}
		}
	}
	fmt.Println(len(days))
	for i := range days {
		fmt.Println(days[i][0], days[i][1], days[i][2])
	}
}
