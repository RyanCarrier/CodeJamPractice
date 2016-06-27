package main

import "fmt"

var tests, n int

var line string
var word1, word2 string

var data map[string]int

var words1 []string
var words2 []string
var freq []int

func main() {
	fmt.Scan(&tests)
	for test := 1; test <= tests; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	words1 = []string{}
	words2 = []string{}
	total := 0
	fmt.Scan(&n)
	for a := 0; a < n; a++ {
		fmt.Scan(&word1)
		words1 = append(words1, word1)
		fmt.Scan(&word2)
		words2 = append(words2, word2)
	}
	for i := 0; i < n; i++ {
		if existExcept(1, words1[i], i) >= 0 {
			if existExcept(2, words2[i], i) >= 0 {
				//fmt.Println(1, words1[i], words2[i])
				total++
			}
		} else if found := existExcept(2, words1[i], -1); found >= 0 {
			if existExcept(2, words2[i], found) >= 0 {
				if n > 2 {
					//fmt.Println(2, words1[i], words2[i])
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func existExcept(which int, find string, except int) int {
	if which == 1 {
		for i, word := range words1 {
			if i == except {
				continue
			}
			if word == find {
				return i
			}
		}
	} else {
		for i, word := range words2 {
			if i == except {
				continue
			}
			if word == find {
				return i
			}
		}
	}
	return -1
}
