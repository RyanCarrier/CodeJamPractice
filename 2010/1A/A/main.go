package main

import "fmt"

//T is tests
var T int //tests
var line string

//N is board size
var N int //board size
//K is rows to win
var K int //in a row to win
var R bool
var B bool

var grid [][]int

func main() {
	fmt.Scan(&T)
	fmt.Println(T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &K)
		/*	fmt.Print(N)
			fmt.Print(" ")
			fmt.Println(K)*/
		grid = make([][]int, N)
		for j := 0; j < N; j++ {
			grid[j] = make([]int, N)
			fmt.Scanln(&line)
			index := N - 1
			for jj := N - 1; jj >= 0; jj-- {
				switch line[jj] {
				case '.':
					continue
				case 'R':
					grid[j][index] = 1
					index--
				case 'B':
					grid[j][index] = -1
					index--
				}
			}
		}
		check()
		var winner string
		switch {
		case R && B:
			winner = "Both"
		case R:
			winner = "Red"
		case B:
			winner = "Blue"
		default:
			winner = "Neither"
		}
		fmt.Printf("Case #%d: %s\n", i, winner)
	}
}

func check() (bool, bool) {
	var row []int

	for i := range grid {
		//Horizontal
		HR, HB := checkSet(grid[i])
		//Verticle
		row = make([]int, N)
		for j := 0; j < N; j++ {
			row[j] = grid[j][i]
		}
		VR, VB := checkSet(row)
		//DiagRight

		var DRR, DRB bool
		if i+K-1 < N {
			for j := 0; j+K-1 < N; j++ {
				TR, TB := checkSet(diagRight(i, j))
				DRR = DRR || TR
				DRB = DRB || TB
			}
		}

		//DiagLeft
		var DLR, DLB bool
		if i-K+1 < N {
			for j := N - 1; j-K+1 >= 0; j-- {
				TR, TB := checkSet(diagLeft(i, j))
				DLR = DLR || TR
				DLB = DLB || TB
			}
		}
		R = HR || VR || DRR || DLR
		B = HB || VB || DRB || DLB
	}

	return R, B
}

func checkSet(set []int) (bool, bool) {
	R, B := false, false
	old := 0
	combo := 0
	for _, v := range set {
		if v == 0 {
			continue
		}
		if v == old {
			combo++
		} else {
			if combo >= K {
				if old == 1 {
					R = true
				} else {
					B = true
				}
			}
			old = v
			combo = 0
		}
	}
	if combo >= K {
		if old == 1 {
			R = true
		} else {
			B = true
		}
	}
	return R, B
}

func diagRight(i, j int) []int {
	var max int
	if i < j {
		max = N - j
	} else {
		max = N - i
	}
	row := make([]int, max)
	for k := 0; k < max; k++ {
		//fmt.Printf("%d, %d\n", i+k, j+k)
		row[k] = grid[i+k][j+k]
	}
	return row
}
func diagLeft(i, j int) []int {
	var max int
	if i < N-j {
		max = j
	} else {
		max = N - i
	}
	row := make([]int, max)
	for k := 0; k < max; k++ {
		//fmt.Printf("%d, %d\n", i+k, j-k)
		row[k] = grid[i+k][j-k]
	}
	return row
}
