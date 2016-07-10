package main

import (
	"bufio"
	"fmt"
	"os"
)

//T test cases
var T int
var bio = bufio.NewReader(os.Stdin)

func main() {
	fmt.Fscan(bio, &T)
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	var line string
	var N int
	fmt.Fscan(bio, &N)
	bestName := ""
	mostUnique := -1

	bio.ReadLine()
	for i := 0; i < N; i++ {
		b, _, _ := bio.ReadLine()
		line = string(b)
		bestName, mostUnique = unique(line, bestName, mostUnique)
	}
	fmt.Println(bestName)
}

func unique(line, bestName string, mostUnique int) (string, int) {
	if mostUnique > len(line) {
		return bestName, mostUnique
	}
	uniqueList := make([]bool, 26)
	for _, a := range line {
		if a == ' ' {
			continue
		}
		uniqueList[a-65] = true
	}
	temp := total(uniqueList)
	if temp == mostUnique {
		if alphFirst(bestName, line) {
			return bestName, mostUnique
		}
		return line, mostUnique
	}
	if temp > mostUnique {
		bestName = line
		mostUnique = temp
	}
	return bestName, mostUnique
}

func total(t []bool) int {
	i := 0
	for _, j := range t {
		if j {
			i++
		}
	}
	return i
}

func alphFirst(a, b string) bool {
	aa := len(a)
	bb := len(b)
	var min int
	if aa < bb {
		min = aa
	} else {
		min = bb
	}
	for i := 0; i < min; i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] < b[i] {
			return true
		}
		return false
	}
	return true
}
