package main

import (
	"fmt"

	ioh "github.com/RyanCarrier/CodeJamPractice/iohelper"
)

//T test cases
var T int

var modulo = int64(1000000007)

//var io = iohelper
func main() {
	T = ioh.GetInt()
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	A, B, N, K := ioh.GetInt64(), ioh.GetInt64(), ioh.GetInt64(), ioh.GetInt64()
	fmt.Println(A, B, N, K)
	total := int64(0)

	ilist := make([]int64, N+1)
	jlist := make([]int64, N+1)
	for i := int64(1); i <= N; i++ {
		ilist[i] = pow(i, A)
		jlist[i] = pow(i, B)
	}
	fmt.Println(ilist[:4], jlist[:4])
	for i := int64(1); i <= N; i++ {
		for j := int64(1); j <= N; j++ {
			if i == j {
				continue
			}
			if (ilist[i]+jlist[j])%K == 0 {
				//	fmt.Println(i, j)
				total++
			}
		}
	}
	fmt.Println(total % modulo)
}

func pow(a, b int64) int64 {
	total := int64(1)
	for i := int64(0); i < b; i++ {
		total *= a
	}
	return total
	//return int64(math.Pow(float64(a), float64(b)))
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
