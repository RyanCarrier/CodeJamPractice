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
	var R, C int
	fmt.Scan(&R, &C)
	if R < 3 || C < 3 {
		fmt.Println(0)
		return
	}
	//imax, jmax := R-1, C-1
	H := make([][]int, R)
	W := make([][]int, R)
	for i := 0; i < R; i++ {
		H[i] = make([]int, C)
		W[i] = make([]int, C)
		for j := 0; j < C; j++ {
			fmt.Scan(&H[i][j])
			if i > 0 && i < R-1 && j > 0 && j < C-1 {
				W[i][j] = 1001

			} else {
				W[i][j] = H[i][j]
			}
		}
	}

	for i := 1; i < R-1; i++ {
		for j := 1; j < C-1; j++ {
			//	W[i][j] = max()
		}
	}
	//	copy(W, H)
	//start away from edges

	//change := true
	//for change {
	//change = false
	m := min(R, C)
	//fmt.Println(m)
	for mm := 0; mm < m; mm++ {
		topi := 1
		boti := R - 2
		leftj := 1
		rightj := C - 2
		for topi <= boti && leftj <= rightj {
			//TOP ROW
			i := topi
			for j := leftj; j <= rightj; j++ {
				W[i][j] = minSlice(W[i][j], W[i-1][j], W[i+1][j], W[i][j-1], W[i][j+1])
				if W[i][j] < H[i][j] {
					W[i][j] = H[i][j]
				}
			}

			//BOTROW
			i = boti
			for j := leftj; j <= rightj; j++ {
				W[i][j] = minSlice(W[i][j], W[i-1][j], W[i+1][j], W[i][j-1], W[i][j+1])
				if W[i][j] < H[i][j] {
					W[i][j] = H[i][j]
				}
			}

			//LEFT COL
			j := leftj
			for i := topi + 1; i < boti; i++ {
				W[i][j] = minSlice(W[i][j], W[i-1][j], W[i+1][j], W[i][j-1], W[i][j+1])
				if W[i][j] < H[i][j] {
					W[i][j] = H[i][j]
				}
			}

			//RIGHT COL
			j = rightj
			for i := topi + 1; i < boti; i++ {
				W[i][j] = minSlice(W[i][j], W[i-1][j], W[i+1][j], W[i][j-1], W[i][j+1])
				if W[i][j] < H[i][j] {
					W[i][j] = H[i][j]
				}
			}
			topi++
			boti--
			leftj++
			rightj--
		}
	}

	//	}
	/*
		fmt.Println("")
		for i := 0; i < R; i++ {
			fmt.Println(H[i])
		}
		fmt.Println("")
		for i := 0; i < R; i++ {
			fmt.Println(W[i])
		}*/
	total := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			total += W[i][j] - H[i][j]
		}
	}
	fmt.Println(total)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSlice(a ...int) int {
	m := 1002
	for _, b := range a {
		m = min(m, b)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlice(a ...int) int {
	m := -1
	for _, b := range a {
		m = max(m, b)
	}
	return m
}
