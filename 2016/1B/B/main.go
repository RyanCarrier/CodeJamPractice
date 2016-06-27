package main

import "fmt"

var tests, n int
var word1, word2 string
var fword1, fword2 string
var line string

func main() {
	fmt.Scan(&tests)
	for test := 1; test <= tests; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	pre := 0 //0 equal, 1 is bigger, 2 is bigger
	indexequal := -1
	fword1, fword2 = "", ""
	fmt.Scan(&word1)
	fmt.Scan(&word2)
	for i := range word1 {
		if word1[i] != '?' && word2[i] != '?' {
			indexequal = i
			if word1[i] == word2[i] {
				pre = 0
			} else {
				if word1[i] > word2[i] {
					pre = 1
				} else {
					pre = 2
				}
			}
			break
		}
	}
	fmt.Println(indexequal)
	for i := range word1 {
		switch {
		case word1[i] != '?' && word2[i] != '?':
			fword1 += string(word1[i])
			fword2 += string(word2[i])
		case word1[i] == '?' && word2[i] == '?':
			switch pre {
			case 0:
				fword1 += "0"
				fword2 += "0"
			case 1:
				fword1 += "0"
				if i > indexequal {
					fword2 += "9"
				} else {
					fword2 += "0"
				}
			case 2:

				fword2 += "0"
				if i > indexequal {
					fword1 += "9"
				} else {
					fword1 += "0"
				}
			}
		case word1[i] == '?':
			fword2 += string(word2[i])
			switch pre {
			case 0:
				fword1 += string(word2[i])
			case 1:
				fword1 += "0"
			case 2:
				if i > indexequal {
					fword1 += "9"
				} else {
					fword1 += "0"
				}
			}
		case word2[i] == '?':
			fword1 += string(word1[i])
			switch pre {
			case 0:
				fword2 += string(word1[i])
			case 1:
				if i > indexequal {
					fword2 += "9"
				} else {
					fword2 += "0"
				}
			case 2:
				fword2 += "0"
			}
		}
	}
	fmt.Println(fword1, fword2)
}
