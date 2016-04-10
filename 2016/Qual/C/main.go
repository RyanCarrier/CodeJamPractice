package main

import (
	"fmt"
	"math"
	"strconv"
)

//T tests
var T int

//N is number seed
var N int

//J comment
var J int
var line string
var ints []bool

var stack []bool

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		fmt.Scan(&N, &J)
		fmt.Printf("Case #%d:\n", test)
		eval()
	}
}

func eval() {
	primeList := make([]int, 9)
	beginString := "1"
	for i := 0; i < N-1; i++ {
		beginString += "0"
	}
	begin, _ := strconv.ParseInt(beginString, 2, 64)
	found := 0
	//fmt.Println(begin, beginString)
	//Prevent inf loop into crash incase I fucked up
	for i := begin; found < J && i < math.MaxInt64; i++ {
		bin := strconv.FormatInt(int64(i), 2)
		if len(bin) != N {
			continue
		}
		//Golang doens't produce leading zeros, we know len is N
		if bin[N-1] == '0' {
			continue
		}
		//fmt.Println(bin)
		p := false
		for base := 2; base <= 10; base++ {
			parsed, _ := strconv.ParseInt(bin, base, 64)
			baseNum := prime(int(parsed))
			if baseNum == 0 {
				p = true
				break
			}
			primeList[base-2] = baseNum
		}
		if !p {
			//fmt.Println(finalList, found)
			fmt.Print("\n", bin)
			for _, prim := range primeList {
				fmt.Print(" ", prim)
			}
			found++
		}
	}
}

func prime(n int) int {
	sr := int(math.Ceil(math.Sqrt(float64(int64(n)))))
	for i := 2; i < sr; i++ {
		if n%i == 0 {
			return i
		}
	}
	return 0
}
