package main

import "fmt"

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
	var M, N int
	fmt.Scan(&M, &N)
	K := make([]int, N)   //max level
	L := make([]int, N)   //current level
	A := make([][]int, N) //attack power at level n
	C := make([][]int, N) //cost to progress from level n

	for i := 0; i < N; i++ {
		fmt.Scan(&K[i], &L[i])
		A[i] = make([]int, K[i])
		for k := 0; k < K[i]; k++ {
			fmt.Scan(&A[i][k])
		}
		C[i] = make([]int, K[i]-1)
		for k := 0; k < K[i]-1; k++ {
			fmt.Scan(&C[i][k])
		}
	}
	/*
		fmt.Println(M)
		fmt.Println(N)
		fmt.Println(K)
		fmt.Println(L)
		fmt.Println(A)
	/	fmt.Println(C)*/
	p := try(M, N, K, L, A, C)
	fmt.Println(p)

}

func try(M, N int, K, L []int, A, C [][]int) int {

	possible := make([]bool, N)
	any := false
	for n := 0; n < N; n++ {
		if L[n] < K[n] && M > C[n][L[n]-1] {
			possible[n] = true
			any = true
		}
	}
	if !any || M == 0 {
		total := 0
		for n := 0; n < N; n++ {
			total += A[n][L[n]-1]
		}
		return total
	}
	bestPower := 0
	for n := 0; n < N; n++ {
		if possible[n] {
			L[n]++
			p := try(M-C[n][L[n]-2], N, K, L, A, C)
			L[n]--
			if p > bestPower {
				bestPower = p
			}
		}

	}
	return bestPower
}
