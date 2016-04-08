package main

import (
	"fmt"
	"strings"
)

//T Tests
var T int

//N number of lines of dirs you have
var N int

//M number of lines of dirs you must make
var M int
var total int
var line string
var root *Dir

//Dir is a dir
type Dir struct {
	Name string
	Dirs map[string]*Dir
}

func newDir(Name string) *Dir {
	return &Dir{Name, map[string]*Dir{}}
}

func (d *Dir) add(these []string, sum bool) {
	if d.Dirs[these[0]] == nil {
		if sum {
			total++
		}
		d.Dirs[these[0]] = newDir(these[0])
	}
	if len(these) > 1 {
		d.Dirs[these[0]].add(these[1:], sum)
	}
}

func (d *Dir) ezAdd(sum bool) {
	fmt.Scan(&line)
	//[1:], remove first /
	d.add(strings.Split(line[1:], "/"), sum)
}

func main() {
	fmt.Scan(&T)
	for test := 1; test <= T; test++ {
		total = 0
		root = newDir("root")
		fmt.Scan(&N, &M)
		for n := 0; n < N; n++ {
			root.ezAdd(false)
		}
		for m := 0; m < M; m++ {
			root.ezAdd(true)
		}
		fmt.Printf("Case #%d: %d\n", test, total)
	}
}
