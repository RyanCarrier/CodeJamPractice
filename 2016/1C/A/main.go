package main

import "fmt"

var tests, n int

var line string
var largest, largest2 int
var largesti, largest2i int
var parties []int
var total int

func main() {
	fmt.Scan(&tests)
	for test := 1; test <= tests; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
		fmt.Println()
	}
}

func solve() {
	var in int
	fmt.Scan(&n)
	parties = make([]int, n)
	total = 0
	largest = 0
	largest2 = 0
	for i := range parties {
		fmt.Scan(&in)
		parties[i] = in
		total += in
		/*
			if in == largest {
				largest2 = in
				largest2i = i
			}
			if in > largest {
				largest2i = largesti
				largest2 = largest
				largesti = i
				largest = in
			}*/
	}
	getNewLargest()
	MinN := (total / 2)
	if total%2 == 1 {
		print(largesti, -1)
		parties[largesti] = largest - 1
	}
	for iteration := 0; iteration < MinN-1; iteration++ {
		getNewLargest()
		doit()
	}
	getNewLargest()
	for largest > 0 && largest2 > 0 {
		getNewLargest()
		doit()
	}

	//stage 1

}

func notEmpty() bool {
	for _, num := range parties {
		if num > 0 {
			return true
		}
	}
	return false
}

func doit() {
	//fmt.Println(largest, largest2)
	if largest > largest2 && largest > 1 {
		print(largesti, largesti)
		parties[largesti] = largest - 2
		if parties[largesti] > largest2 {
			largest -= 2
			return
		}
	} else {
		print(largesti, largest2i)
		parties[largesti] = largest - 1
		largest--
		parties[largest2i] = largest2 - 1
		largest2--
	}
}

func print(a, b int) {
	A := int('A') + a
	if b != -1 {
		B := int('A') + b
		fmt.Print(string(rune(A))+string(rune(B)), " ")
		return
	}
	fmt.Print(string(rune(A)), " ")
}

func getNewLargest() {
	largest = 0
	largesti = -1
	largest2 = 0
	largest2i = -1
	for i, num := range parties {
		if num == largest {
			largest2 = num
			largest2i = i
		}
		if num > largest {
			largest2i = largesti
			largest2 = largest
			largesti = i
			largest = num
		}

	}
}
