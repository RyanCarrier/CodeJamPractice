package main

import "fmt"

//T tests
var T int

//K is the amount of tiles in orig
var K int

//C is the complexity of what you have ot work wiht, how many iterations are ran.
var C int

//S is the amount of tiles you are allowed to check
var S int

//J comment
var J int
var line string
var ints []bool

var stack []bool

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Scan(&K, &C, &S)
		fmt.Printf("Case #%d:", test)
		results := eval()
		//fmt.Println(K, C, S, results)
		if len(results) > S {
			fmt.Print("IMPOSSIBLE")
		} else {
			for _, r := range results {
				fmt.Print(" ", r)
			}
		}
		fmt.Println()
	}
}

func eval() []int {
	if K == 1 && S > 0 {
		return []int{1}
	}
	if C == 1 {
		result := make([]int, K)
		for i := 0; i < K; i++ {
			result[i] = i + 1
		}
		return result
	}
	result := []int{}
	if K < C-1 {
		return []int{}
	}
	for i := 0; i < K-(C-1); i++ {
		col := i * K
		row := i + 2
		result = append(result, col+row)
	}
	return result
}
